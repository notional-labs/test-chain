package keeper

import (
	"github.com/notional-labs/test-chain2/x/testchain2/types"
)

var _ types.QueryServer = Keeper{}
