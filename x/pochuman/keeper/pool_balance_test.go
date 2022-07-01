package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/VigorousDeveloper/poc-human/testutil/keeper"
	"github.com/VigorousDeveloper/poc-human/testutil/nullify"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/keeper"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPoolBalance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PoolBalance {
	items := make([]types.PoolBalance, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPoolBalance(ctx, items[i])
	}
	return items
}

func TestPoolBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.PochumanKeeper(t)
	items := createNPoolBalance(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoolBalance(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPoolBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.PochumanKeeper(t)
	items := createNPoolBalance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoolBalance(ctx,
			item.Index,
		)
		_, found := keeper.GetPoolBalance(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPoolBalanceGetAll(t *testing.T) {
	keeper, ctx := keepertest.PochumanKeeper(t)
	items := createNPoolBalance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPoolBalance(ctx)),
	)
}
