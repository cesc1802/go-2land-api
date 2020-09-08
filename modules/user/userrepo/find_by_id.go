package userrepo

import (
	"context"
	"go-rest-api/modules/user/usermodel"
)

type FindUserByIdStorage interface {
	FindUserById(ctx context.Context, userId string) (usermodel.User, error)
}

type findUserByIdStorage struct {
	store FindUserByIdStorage
}

func NewFindUserById(storage FindUserByIdStorage) *findUserByIdStorage {
	return &findUserByIdStorage{
		store: storage,
	}
}

func (fbid *findUserByIdStorage) FindUserById(ctx context.Context, userId string) (usermodel.User, error) {
	return fbid.store.FindUserById(ctx, userId)
}
