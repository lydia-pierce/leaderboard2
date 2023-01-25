package leaderboard_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "leaderboard2/testutil/keeper"
	"leaderboard2/testutil/nullify"
	"leaderboard2/x/leaderboard"
	"leaderboard2/x/leaderboard/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PlayerInfoList: []types.PlayerInfo{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LeaderboardKeeper(t)
	leaderboard.InitGenesis(ctx, *k, genesisState)
	got := leaderboard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PlayerInfoList, got.PlayerInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
