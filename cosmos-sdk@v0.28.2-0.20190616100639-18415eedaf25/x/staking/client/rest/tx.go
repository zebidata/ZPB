package rest

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"

	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/delegations",
		postDelegationsHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/unbonding_delegations",
		postUnbondingDelegationsHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/redelegations",
		postRedelegationsHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/staking/validators/{validatorAddr}/stake",
		poststakingsHandlerFn(cliCtx),
	).Methods("POST")
}

type (
	// DelegateRequest defines the properties of a delegation request's body.
	DelegateRequest struct {
		BaseReq          rest.BaseReq   `json:"base_req"`
		DelegatorAddress sdk.AccAddress `json:"delegator_address"` // in bech32
		ValidatorAddress sdk.ValAddress `json:"validator_address"` // in bech32
		Amount           sdk.Coin       `json:"amount"`
	}

	// RedelegateRequest defines the properties of a redelegate request's body.
	RedelegateRequest struct {
		BaseReq             rest.BaseReq   `json:"base_req"`
		DelegatorAddress    sdk.AccAddress `json:"delegator_address"`     // in bech32
		ValidatorSrcAddress sdk.ValAddress `json:"validator_src_address"` // in bech32
		ValidatorDstAddress sdk.ValAddress `json:"validator_dst_address"` // in bech32
		Amount              sdk.Coin       `json:"amount"`
	}

	// UndelegateRequest defines the properties of a undelegate request's body.
	UndelegateRequest struct {
		BaseReq          rest.BaseReq   `json:"base_req"`
		DelegatorAddress sdk.AccAddress `json:"delegator_address"` // in bech32
		ValidatorAddress sdk.ValAddress `json:"validator_address"` // in bech32
		Amount           sdk.Coin       `json:"amount"`
	}

	// StakeRequest defines the properties of a stake request's body.
	StakingRequest struct {
		BaseReq                 rest.BaseReq `json:"base_req"`
		Commissionrate          string       `json:"commissionrate"`
		Commissionmaxrate       string       `json:"commissionmaxrate"`
		Commissionmaxchangerate string       `json:"commissionmaxchangerate"`
		//DelegatorAddress  						string  `json:"delegator_address"`
		ValidatorAddress  sdk.ValAddress `json:"validatoraddress"`
		Minselfdelegation string         `json:"minselfdelegation"`
		PubKey            string         `json:"pubkey"`
		Moniker           string         `json:"moniker"`
		Identity          string         `json:"identity"`
		Website           string         `json:"website"`
		Details           string         `json:"details"`
		Denom             string         `json:"denom"`
		Amount            string         `json:"amount"`
	}
)

func postDelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DelegateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgDelegate(req.DelegatorAddress, req.ValidatorAddress, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postRedelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RedelegateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgBeginRedelegate(req.DelegatorAddress, req.ValidatorSrcAddress, req.ValidatorDstAddress, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postUnbondingDelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UndelegateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUndelegate(req.DelegatorAddress, req.ValidatorAddress, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func poststakingsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req StakingRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			panic(err)
		}

		var pubkey crypto.PubKey
		pubkey, err = sdk.GetConsPubKeyBech32(req.PubKey)

		var commissions types.CommissionRates

		commissions.Rate, err = sdk.NewDecFromStr(req.Commissionrate)
		commissions.MaxRate, err = sdk.NewDecFromStr(req.Commissionmaxrate)
		commissions.MaxChangeRate, err = sdk.NewDecFromStr(req.Commissionmaxchangerate)

		var description types.Description
		description.Moniker = req.Moniker
		description.Identity = req.Identity
		description.Website = req.Website
		description.Details = req.Details
		//var m_s_delegation sdk.Int
		m_s_delegation, errr := sdk.NewIntFromString(req.Minselfdelegation)
		if errr != true {
			return
		}

		var amount sdk.Coin
		amount.Denom = req.Denom
		abc, errr := sdk.NewIntFromString(req.Amount)
		if errr != true {
			return
		}
		amount.Amount = abc

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateValidator(req.ValidatorAddress, pubkey, amount, description, commissions, m_s_delegation)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if !bytes.Equal(fromAddr, req.ValidatorAddress) {
			//rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own validator address")
			//return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
