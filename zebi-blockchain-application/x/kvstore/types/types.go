package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Data is a struct that contains all the metadata of a data
type Data struct {
	Key    string         `json:"key"`
	Value  string         `json:"value"`
	Sender sdk.AccAddress `json:"sender"`
}

// NewKeyValue returns an empty string
func NewKeyValue() Data {
	return Data{}
}

// implement fmt.Stringer
func (w Data) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Key: %s
Value: %s
Sender: %s`, w.Key, w.Value, w.Sender))
}


// NewSign returns an empty string
func NewSign() string {
	return ""
}
