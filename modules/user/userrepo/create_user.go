package userrepo

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, createUser usermodel.User) (usermodel.SimpleUser, error)
}

type createUserStorage struct {
	uStore CreateUserStorage
}

func NewCreateUserStorage(createUserStore CreateUserStorage) *createUserStorage {
	return &createUserStorage{
		uStore: createUserStore,
	}
}

func (repo *createUserStorage) CreateUser(ctx context.Context, createUser usermodel.CreateUser) (usermodel.SimpleUser, error) {
	user := createUser.ToUser()
	return repo.uStore.CreateUser(ctx, user)
}
