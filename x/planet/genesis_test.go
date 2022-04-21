package planet_test

import (
	"testing"

	keepertest "github.com/klpanda/planet/testutil/keeper"
	"github.com/klpanda/planet/testutil/nullify"
	"github.com/klpanda/planet/x/planet"
	"github.com/klpanda/planet/x/planet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlanetKeeper(t)
	planet.InitGenesis(ctx, *k, genesisState)
	got := planet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
