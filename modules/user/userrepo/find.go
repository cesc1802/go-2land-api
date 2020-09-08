package userrepo

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type FindUserStorage interface {
	FindUserByUsername(ctx context.Context, user usermodel.User) (usermodel.User, error)
}

type findUserStorage struct {
	store FindUserStorage
}

func NewFindUserStorage(store FindUserStorage) *findUserStorage {
	return &findUserStorage{
		store: store,
	}
}

func (u *findUserStorage) FindUserByUsername(ctx context.Context, user usermodel.User) (usermodel.User, error) {
	return u.store.FindUserByUsername(ctx, user)
}
