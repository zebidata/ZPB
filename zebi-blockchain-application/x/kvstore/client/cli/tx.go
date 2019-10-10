package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"zebi-blockchain-application/x/kvstore/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	kvstoreTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "kvstore transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	kvstoreTxCmd.AddCommand(client.PostCommands(
		GetCmdPostKeyValue(cdc),
	)...)

	return kvstoreTxCmd
}

// GetCmdPostKeyValue is the CLI command for sending a key value transaction
func GetCmdPostKeyValue(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post-key-value [key] [value] [flags]",
		Short: "post key value pair of any string length",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			/*coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}*/

			msg := types.NewMsgPostKeyValue(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
