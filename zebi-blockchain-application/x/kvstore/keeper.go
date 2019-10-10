package kvstore

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// PostKeyValue posts key-value string of arbitrary length
func (k Keeper) PostKeyValue(ctx sdk.Context, data Data) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(data.Key), k.cdc.MustMarshalBinaryBare(data))
}

// GetKeyValue - returns the string that key resolves to
func (k Keeper) GetKeyValue(ctx sdk.Context, key string) Data {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(key)) {
		return NewKeyValue()
	}
	bz := store.Get([]byte(key))
	var data Data
	k.cdc.MustUnmarshalBinaryBare(bz, &data)
	return data
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetDataIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
