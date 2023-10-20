package keeper_test

import (
	"testing"

	testkeeper "course-blockchain/testutil/keeper"
	"course-blockchain/x/courseblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CourseblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
