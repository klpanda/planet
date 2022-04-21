package keeper

import (
	"github.com/klpanda/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
