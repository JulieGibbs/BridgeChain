package pochuman

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Keysign tx request
	k.KeysignTxRequest(ctx)

	// Update tx request
	k.UpdateTxRequestByObservationVote(ctx)
}
