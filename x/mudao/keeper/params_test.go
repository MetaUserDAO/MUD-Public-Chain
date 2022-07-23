package keeper_test

import (
	"testing"

	testkeeper "github.com/michaelz/mudao/testutil/keeper"
	"github.com/michaelz/mudao/x/mudao/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MudaoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
