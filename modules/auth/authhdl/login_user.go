package authhdl

import (
	"context"
	"go-rest-api/modules/auth/authmodel"
)

type LoginUserHdl interface {
	Login(ctx context.Context, user authmodel.LoginUser) (string, error)
}

type loginUserHdl struct {
	hdl LoginUserHdl
}

func NewLoginUserHdl(hdl LoginUserHdl) *loginUserHdl {
	return &loginUserHdl{
		hdl: hdl,
	}
}

func (luh *loginUserHdl) Login(ctx context.Context, user authmodel.LoginUser) (string, error) {
	return luh.hdl.Login(ctx, user)
}