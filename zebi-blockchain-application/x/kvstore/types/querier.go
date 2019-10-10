package types

import "strings"

// Query Result Payload for a names query
type QueryResKeyValue []string

// implement fmt.Stringer
func (n QueryResKeyValue) String() string {
	return strings.Join(n[:], "\n")
}