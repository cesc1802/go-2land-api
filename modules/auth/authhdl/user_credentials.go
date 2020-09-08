package authhdl

import (
	"context"
	"go-rest-api/modules/auth/authmodel"
)

type UserCredentialsHdl interface {
	Login(ctx context.Context, userCredentials authmodel.UserCredentials) (authmodel.UserAuthInfo, error)
}

type userCredentialsHdl struct {
	uHandler UserCredentialsHdl
}

func NewUserCredentialsHdl(uHandler UserCredentialsHdl) *userCredentialsHdl {
	return &userCredentialsHdl{
		uHandler: uHandler,
	}
}

func (uch *userCredentialsHdl) Login(ctx context.Context, userCredentials authmodel.UserCredentials) (authmodel.UserAuthInfo, error) {
	return uch.uHandler.Login(ctx, userCredentials)
}
