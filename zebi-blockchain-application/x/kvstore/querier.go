package kvstore

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the kvstore Querier
const (
	QueryKeyValue = "key"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryKeyValue:
			return queryKeyValue(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown kvstore query endpoint")
		}
	}
}

// nolint: unparam
func queryKeyValue(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	data := keeper.GetKeyValue(ctx, path[0])
	value := data.Value
	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not get value")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResKeyValue{value})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
