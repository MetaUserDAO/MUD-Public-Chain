package keeper

import (
	"github.com/michaelz/mudao/x/mudao/types"
)

var _ types.QueryServer = Keeper{}
