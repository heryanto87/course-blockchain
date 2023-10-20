package keeper

import (
	"course-blockchain/x/courseblockchain/types"
)

var _ types.QueryServer = Keeper{}
