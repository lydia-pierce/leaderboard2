package keeper

import (
	"leaderboard2/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
