package mudao_test

import (
	"testing"

	keepertest "github.com/michaelz/mudao/testutil/keeper"
	"github.com/michaelz/mudao/testutil/nullify"
	"github.com/michaelz/mudao/x/mudao"
	"github.com/michaelz/mudao/x/mudao/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MudaoKeeper(t)
	mudao.InitGenesis(ctx, *k, genesisState)
	got := mudao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
