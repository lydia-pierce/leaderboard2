package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "leaderboard2/testutil/keeper"
	"leaderboard2/x/leaderboard/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeaderboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
