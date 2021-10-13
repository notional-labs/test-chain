package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/notional-labs/test-chain2/testutil/keeper"
	"github.com/notional-labs/test-chain2/x/testchain2/keeper"
	"github.com/notional-labs/test-chain2/x/testchain2/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Testchain2Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
