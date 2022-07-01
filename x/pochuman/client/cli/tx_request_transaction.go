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

func CmdRequestTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-transaction [origin-chain] [origin-address] [target-chain] [target-address] [amount] [fee]",
		Short: "Broadcast message request-transaction",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOriginChain := args[0]
			argOriginAddress := args[1]
			argTargetChain := args[2]
			argTargetAddress := args[3]
			argAmount := args[4]
			argFee := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestTransaction(
				clientCtx.GetFromAddress().String(),
				argOriginChain,
				argOriginAddress,
				argTargetChain,
				argTargetAddress,
				argAmount,
				argFee,
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
