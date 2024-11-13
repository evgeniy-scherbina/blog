package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuthor(goCtx context.Context, msg *types.MsgCreateAuthor) (*types.MsgCreateAuthorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var author = types.Author{
		Creator:   msg.Creator,
		Firstname: msg.Firstname,
		Lastname:  msg.Lastname,
	}
	id := k.AppendAuthor(
		ctx,
		author,
	)

	return &types.MsgCreateAuthorResponse{
		Id: id,
	}, nil
}
