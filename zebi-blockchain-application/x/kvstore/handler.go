package kvstore

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "kvstore" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgPostKeyValue:
			return handleMsgPostKeyValue(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized kvstore Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to post key value pair
func handleMsgPostKeyValue(ctx sdk.Context, keeper Keeper, msg MsgPostKeyValue) sdk.Result {
	if msg.Sender.Empty() { // Checks if the the msg sender exist
		return sdk.ErrUnauthorized("Incorrect Sender").Result() // If not, throw an error
	}
	var data Data
	data.Key = msg.Key
	data.Value = msg.Value
	data.Sender = msg.Sender
	keeper.PostKeyValue(ctx, data) // If so, set the value to the key specified in the msg.
	resTags := sdk.NewTags(
		"type", "kvstore",
		"sender", msg.Sender.String(),
		"key", msg.Key,
		"value", msg.Value,
	)
	return sdk.Result{
		Tags: resTags,
		}                          // return
}
