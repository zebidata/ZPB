package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"zebi-blockchain-application/x/kvstore/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	kvstoreQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the kvstore module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	kvstoreQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetKeyValue(storeKey, cdc),
	)...)
	return kvstoreQueryCmd
}

// GetCmdGetKeyValue queries information about a key
func GetCmdGetKeyValue(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-key-value [key]",
		Short: "gets value existing on blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/key/%s", queryRoute, key), nil)
			if err != nil {
				fmt.Printf("could not get value for key - %s \n", string(key))
				return nil
			}

			var out types.QueryResKeyValue
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
