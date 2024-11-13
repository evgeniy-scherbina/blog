package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuthor(goCtx context.Context, msg *types.MsgCreateAuthor) (*types.MsgCreateAuthorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateAuthorResponse{}, nil
}
