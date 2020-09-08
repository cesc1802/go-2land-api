package userhdl

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type CreateUserHdl interface {
	CreateUser(ctx context.Context, createUser usermodel.CreateUser) (usermodel.SimpleUser, error)
}

type createUserHdl struct {
	handler CreateUserHdl
}

func NewCreateUserHdl(handler CreateUserHdl) *createUserHdl {
	return &createUserHdl{
		handler: handler,
	}
}

func (hdl *createUserHdl) CreateUser(ctx context.Context, createUser usermodel.CreateUser) (usermodel.SimpleUser, error) {
	return hdl.handler.CreateUser(ctx, createUser)
}
