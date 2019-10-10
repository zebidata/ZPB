package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgPostKeyValue defines a MsgPostKeyValue message
type MsgPostKeyValue struct {
	Key    string         `json:"key"`
	Value  string		  `json:"value"`
	Sender sdk.AccAddress `json:"sender"`
}

// NewMsgPostKeyValue is a constructor function for NewMsgPostKeyValue
func NewMsgPostKeyValue(key string, value string, sender sdk.AccAddress) MsgPostKeyValue {
	return MsgPostKeyValue{
		Key:    key,
		Value:  value,
		Sender: sender,
	}
}

// Route should return the name of the module
func (msg MsgPostKeyValue) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPostKeyValue) Type() string { return "post_kv" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPostKeyValue) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	if len(msg.Key) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Key Value pair cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPostKeyValue) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPostKeyValue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}



// MsgPostSign defines a MsgPostSign message
type MsgPostSign struct {
	Sender sdk.AccAddress `json:"sender"`
}

// NewMsgPostSign is a constructor function for NewMsgPostSign
func NewMsgPostSign( sender sdk.AccAddress) MsgPostSign {
	return MsgPostSign{
		
		Sender: sender,
	}
}

// Route should return the name of the module
func (msg MsgPostSign) Route() string { return RouterKey }


// Type should return the action
func (msg MsgPostSign) Type() string { return "post_sign" }


// ValidateBasic runs stateless checks on the message
func (msg MsgPostSign) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPostSign) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPostSign) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
