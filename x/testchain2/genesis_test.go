package testchain2_test

import (
	"testing"

	keepertest "github.com/notional-labs/test-chain2/testutil/keeper"
	"github.com/notional-labs/test-chain2/testutil/nullify"
	"github.com/notional-labs/test-chain2/x/testchain2"
	"github.com/notional-labs/test-chain2/x/testchain2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Testchain2Keeper(t)
	testchain2.InitGenesis(ctx, *k, genesisState)
	got := testchain2.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
