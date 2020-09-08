package userrepo

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type FindByUsernameStorage interface {
	FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error)
}

type findByUsernameStorage struct {
	uStore FindByUsernameStorage
}

func NewFindByUsernameStorage(uStore FindByUsernameStorage) *findByUsernameStorage {
	return &findByUsernameStorage{
		uStore: uStore,
	}
}

func (repo *findByUsernameStorage) FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error) {
	return repo.uStore.FindUserByUsername(ctx, userName)
}
