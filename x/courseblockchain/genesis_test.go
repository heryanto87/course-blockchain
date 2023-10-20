package courseblockchain_test

import (
	"testing"

	keepertest "course-blockchain/testutil/keeper"
	"course-blockchain/testutil/nullify"
	"course-blockchain/x/courseblockchain"
	"course-blockchain/x/courseblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CourseblockchainKeeper(t)
	courseblockchain.InitGenesis(ctx, *k, genesisState)
	got := courseblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
