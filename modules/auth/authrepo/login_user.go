package authrepo

import (
	"context"
	"fmt"
	"go-rest-api/modules/auth/authmodel"
	"go-rest-api/modules/user/usermodel"
	"go-rest-api/modules/user/userrepo"
)

type LoginUserStorage interface {
}

type loginUserStorage struct {
	userStore  userrepo.FindUserStorage
	loginStore LoginUserStorage
}

func NewLoginUserStorage(userStore userrepo.FindUserStorage, loginStore LoginUserStorage) *loginUserStorage {
	return &loginUserStorage{
		userStore:  userStore,
		loginStore: loginStore,
	}
}

func (lu *loginUserStorage) FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error) {
	return lu.userStore.FindUserByUsername(ctx, userName)
}

func (lu *loginUserStorage) Login(ctx context.Context, inputUser authmodel.LoginUser) (string, error) {
	var user *usermodel.User
	var err error
	user, err = lu.FindUserByUsername(ctx, inputUser.Username)

	fmt.Println(user)
	fmt.Println(err)
	return "abc", nil
}
