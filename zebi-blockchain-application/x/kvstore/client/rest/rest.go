package rest

import (
	"fmt"
    "net/http"
    "io/ioutil"
	"encoding/json"
	"errors"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"zebi-blockchain-application/x/kvstore/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	cliKeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	authType "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	bip39 "github.com/cosmos/go-bip39"
//	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/gorilla/mux"
)

const (
	restKey  = "key"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeData string) {
	r.HandleFunc(fmt.Sprintf("/%s/key", storeData), postKeyValueHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/key/{%s}", storeData, restKey), getKeyValueHandler(cliCtx, storeData)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/signTx", storeData), postSignHandler( cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/Acc"), postAddHandler( cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/ConPub"), postConPubHandler( cliCtx)).Methods("POST")
	
}

type postKeyValueReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Key     string       `json:"key"`
	Value   string       `json:"value"`
	Sender  string       `json:"sender"`
}

func postKeyValueHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req postKeyValueReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Sender)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		/*
			coins, err := sdk.ParseCoins(req.Amount)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}*/

		// create the message
		msg := types.NewMsgPostKeyValue(req.Key, req.Value, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func getKeyValueHandler(cliCtx context.CLIContext, storeData string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restKey]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/key/%s", storeData, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type postSignReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Password string 	 `json:"pass"`
	ArmorPrivKeys   string 	 		 `json:"armorprivkeys"`
}


func postSignHandler( cliCtx context.CLIContext)  http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		var stdTx authType.StdTx

	
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		bodyTemp:=body

		err = cliCtx.Codec.UnmarshalJSON(body, &stdTx)
		if err != nil {
			panic(err)
		}
		var extraParam postSignReq

	
		err = json.Unmarshal(bodyTemp, &extraParam)
		if err != nil {
		panic(err)
		}
	
		//var extraParamPriv privKeysStruct
		privArmor:= extraParam.ArmorPrivKeys
		priv, err := mintkey.UnarmorDecryptPrivKey(privArmor, extraParam.Password)
	
		GasAdjustmentValue, _ := rest.ParseFloat64OrReturnBadRequest(w, extraParam.BaseReq.GasAdjustment, flags.DefaultGasAdjustment)
		
		simAndExec, GasValue, errGas := flags.ParseGas(extraParam.BaseReq.Gas)
		if errGas != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errGas.Error())
			return
		}
	
		gasPrice:=extraParam.BaseReq.GasPrices

		fees := stdTx.Fee.Amount
		if !gasPrice.IsZero() {
			if !fees.IsZero() {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot provide both fees and gas prices").Error())
			return
		}

		glDec := sdk.NewDec(int64(GasValue))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		fees = make(sdk.Coins, len(gasPrice))
		for i, gp := range gasPrice {
			fee := gp.Amount.Mul(glDec)
			fees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}

	}
		cdc := codec.New()

		txBldr := authType.NewTxBuilderFromJSON(
			utils.GetTxEncoder(cdc),
			extraParam.BaseReq.AccountNumber,
			extraParam.BaseReq.Sequence,
			GasValue,
			GasAdjustmentValue,
			extraParam.BaseReq.Simulate,
			extraParam.BaseReq.ChainID,
			extraParam.BaseReq.Memo,
			stdTx.Fee.Amount,
			gasPrice,	
		)
		var newTx authType.StdTx
		appendSig := false
		offline := false
		generateSignatureOnly:= false
		
		
	stdTxNew := authType.NewStdTx(stdTx.GetMsgs(),authType.NewStdFee(txBldr.Gas(),fees), stdTx.GetSignatures(),stdTx.GetMemo()) 


	if extraParam.BaseReq.Simulate || simAndExec {
		if GasAdjustmentValue < 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.New("invalid gas adjustment").Error())
			return
		}
		
		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, stdTxNew.GetMsgs())
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if extraParam.BaseReq.Simulate {
			rest.WriteSimulationResponse(w, cliCtx.Codec, txBldr.Gas())
			return
		}
	}
		
		newTx, err = utils.SignStdTxPass(txBldr, cliCtx,extraParam.BaseReq.From, priv, stdTxNew, appendSig, offline)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//create signature json
		json, err := getSignatureJSON( cliCtx, newTx, cliCtx.Indent, generateSignatureOnly)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
			//method call to send response
			rest.PostProcessResponse(w, cliCtx, json)


	}
}


func getSignatureJSON( cliCtx context.CLIContext,newTx authType.StdTx, indent, generateSignatureOnly bool) ([]byte, error) {
	switch generateSignatureOnly {
	case true:
		switch indent {
		case true:
			return  cliCtx.Codec.MarshalJSONIndent(newTx.Signatures[0], "", "  ")

		default:
			return  cliCtx.Codec.MarshalJSON(newTx.Signatures[0])
		}
	default:
		switch indent {
		case true:
			return  cliCtx.Codec.MarshalJSONIndent(newTx, "", "  ")

		default:
			return  cliCtx.Codec.MarshalJSON(newTx)
		}
	}
}


type postAddReq struct {
	AddName string 	 `json:"name"`
	AddType string 	 `json:"type"`
	AddPass string 	 `json:"pass"`
	AddMnemonic string 	 `json:"mnemonic"`

}


func postAddHandler( cliCtx context.CLIContext)  http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		
	var kb keys.Keybase
	var err error
	var encryptPassword string
	
	//body Read
	body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		bodyTemp:=body
	
	var extraParam postAddReq
	
	err = json.Unmarshal(bodyTemp, &extraParam)
		if err != nil {
		panic(err)
		}
	var  bip39Passphrase string
	//Should be fixed i.e. non-immutable 
	name := extraParam.AddName
	
	//type:"add"/"recover"
	addType := extraParam.AddType
	
	//pass should be encoded
	//decoding to be done
	pass := extraParam.AddPass

	mnemonic:= extraParam.AddMnemonic

//	interactive := false //viper.GetBool(flagInteractive)
	showMnemonic := true

		kb, err = cliKeys.NewKeyBaseFromHomeFlag()
		if err != nil {
			panic(err)
		}

		//to throw error in response body
	encryptPassword=pass
	
	//default values for account and index

	account :=uint32(0); // uint32(viper.GetInt(flagAccount))
	
	index := uint32(0); //uint32(viper.GetInt(flagIndex))

	if addType=="recover" || addType=="add" {

		if addType=="add" {
			_, err = kb.Get(name)
			if err == nil {
				// account exists, ask for user confirmation
				panic(errors.New("Account with this name already exists"));
			}
		}
	if addType=="recover" {
		mnemonic=strings.TrimSpace(mnemonic)
		if !bip39.IsMnemonicValid(mnemonic) {
			
			panic(errors.New("invalid mnemonic")) ;
		}
	}
	
	} else {
		panic(errors.New("Type must be either add/recover"));
	}

	if len(mnemonic) == 0 && addType=="add" {
		// read entropy seed straight from crypto.Rand and convert to mnemonic
		entropySeed, err := bip39.NewEntropy(256)
		if err != nil {
			panic(err)
		}

		mnemonic, err = bip39.NewMnemonic(entropySeed[:])
		if err != nil {
			panic(err)
		}
	}

	info, err := kb.CreateAccount(name, mnemonic, bip39Passphrase, encryptPassword, account, index)
	if err != nil {
		panic(err)
	}

	// Recover key from seed passphrase
	if addType=="recover" {
		// Hide mnemonic from output
		showMnemonic = false
		mnemonic = ""
	}
			json, err:= printCreate(info, showMnemonic, mnemonic,cliCtx,encryptPassword )
			//method call to send response
			rest.PostProcessResponse(w, cliCtx, json)
	}
}

func printCreate(info keys.Info, showMnemonic bool, mnemonic string,cliCtx context.CLIContext ,encryptPassword string) ([]byte, error){
	
	out, err := keys.Bech32KeyPrivOutput(info,encryptPassword)
	
	if err != nil {
		panic(err)
	}

	if showMnemonic {
		out.Mnemonic = mnemonic
	}

	var jsonString []byte
	if cliCtx.Indent {
		jsonString, err =  cliCtx.Codec.MarshalJSONIndent(out, "", "  ")
	} else {
		jsonString, err =  cliCtx.Codec.MarshalJSON(out)
	}

	if err != nil {
		panic(err)
	}
return jsonString,err
}



type postConvReq struct {

	Password 		string 	 		 `json:"pass"`
	ArmorPrivKeys   string 	 		 `json:"armorprivkeys"`

	
}
//PubOutput for output
type PubOutput struct {
	PubKey    string                 `json:"pubkey"`
}

func postConPubHandler( cliCtx context.CLIContext)  http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		var conv postConvReq

		// body done open
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &conv)
		if err != nil {
			panic(err)
		}

		//var extraParamPriv privKeysStruct
		privArmor:= conv.ArmorPrivKeys
		priv, errPriv := mintkey.UnarmorDecryptPrivKey(privArmor, conv.Password)
		if errPriv != nil {
			panic(errPriv)
		}
	
		pubkey := priv.PubKey()
		
		bechPub,errbech :=sdk.Bech32ifyAccPub(pubkey)
		if errbech != nil {
			panic( errbech)
		}
	
		po := PubOutput{
			PubKey:  bechPub,
		}

		//create signature json
		json, errjson := printPub(po,cliCtx)
		if errjson != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errjson.Error())
			return
		}
			//method call to send response
			rest.PostProcessResponse(w, cliCtx, json)


	}
}


func printPub(pub PubOutput,cliCtx context.CLIContext) ([]byte, error){
	 var err error
	var jsonString []byte
	if cliCtx.Indent {
		jsonString, err =  cliCtx.Codec.MarshalJSONIndent(pub, "", "  ")
	} else {
		jsonString, err =  cliCtx.Codec.MarshalJSON(pub)
	}

	if err != nil {
		panic(err)
	}
return jsonString,err
}


