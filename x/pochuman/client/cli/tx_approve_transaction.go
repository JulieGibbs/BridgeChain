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

func CmdApproveTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-transaction [tx-hash] [success] [signed-key]",
		Short: "Broadcast message approve-transaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxHash := args[0]
			argSuccess := args[1]
			argSignedKey := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveTransaction(
				clientCtx.GetFromAddress().String(),
				argTxHash,
				argSuccess,
				argSignedKey,
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
