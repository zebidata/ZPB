package kvstore

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	DataRecords []Data `json:"data_records"`
}

func NewGenesisState(dataRecords []Data) GenesisState {
	return GenesisState{DataRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.DataRecords {
		if record.Key == "" {
			return fmt.Errorf("Invalid DataRecord: Value: %s. Error: Missing Key", record.Key)
		}
		if record.Value == "" {
			return fmt.Errorf("Invalid DataRecord: Owner: %s. Error: Missing Value", record.Value)
		}
		if record.Sender == nil {
			return fmt.Errorf("Invalid DataRecord: Value: %s. Error: Missing Sender", record.Sender)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		DataRecords: []Data{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.DataRecords {
		keeper.PostKeyValue(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Data
	iterator := k.GetDataIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		var value Data
		value = k.GetKeyValue(ctx, key)
		records = append(records, value)
	}
	return GenesisState{DataRecords: records}
}
