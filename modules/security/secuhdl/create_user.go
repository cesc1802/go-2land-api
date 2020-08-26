package secuhdl

import (
	"context"
	models "go-rest-api/modules/security/secumodel"
)

type CreateUserHdl interface {
	Create(ctx context.Context, user models.Security) error
}

type createUserHdl struct {
	userHdl CreateUserHdl
}

func NewCreateUserHdl(userHdl CreateUserHdl) *createUserHdl {
	return &createUserHdl{
		userHdl: userHdl,
	}
}

func (c *createUserHdl) CreateUser(ctx context.Context, user models.Security) error {
	return c.userHdl.Create(ctx, user)
}
