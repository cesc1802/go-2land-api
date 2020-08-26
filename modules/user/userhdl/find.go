package userhdl

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type FindUserHdl interface {
	FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error)
}

type findUserHdl struct {
	hdl FindUserHdl
}

func NewFindUserHdl(hdl FindUserHdl) *findUserHdl {
	return &findUserHdl{
		hdl: hdl,
	}
}

func (h *findUserHdl) FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error) {
	return h.hdl.FindUserByUsername(ctx, userName)
}
