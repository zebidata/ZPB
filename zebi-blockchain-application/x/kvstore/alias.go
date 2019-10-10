package kvstore

import (
	"zebi-blockchain-application/x/kvstore/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgPostKeyValue 	= types.NewMsgPostKeyValue
	NewMsgPostSign 		= types.NewMsgPostSign
	NewKeyValue         = types.NewKeyValue
	ModuleCdc     		= types.ModuleCdc
	RegisterCodec		= types.RegisterCodec
)

type (
	MsgPostKeyValue     = types.MsgPostKeyValue
	QueryResKeyValue    = types.QueryResKeyValue
	Data           		= types.Data
	MsgPostSign         = types.MsgPostSign
)
