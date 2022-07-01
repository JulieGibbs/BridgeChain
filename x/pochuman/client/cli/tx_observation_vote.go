package cli

import (
	"strconv"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdObservationVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "observation-vote [tx-hash] [chain-name] [from] [to] [amount]",
		Short: "Broadcast message observation-vote",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxHash := args[0]
			argChainName := args[1]
			argFrom := args[2]
			argTo := args[3]
			argAmount := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgObservationVote(
				clientCtx.GetFromAddress().String(),
				argTxHash,
				argChainName,
				argFrom,
				argTo,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
