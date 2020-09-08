package authrepo

import (
	"context"
	common "go-rest-api/common/errors"
	"go-rest-api/utils"
	"time"

	//"fmt"
	"go-rest-api/common/token"
	"go-rest-api/common/token/jwt"
	"go-rest-api/modules/auth/authmodel"
	"go-rest-api/modules/user/usermodel"
	"go-rest-api/modules/user/userrepo"
)

type UserCredentialsStorage interface {
	//Login(ctx context.Context, userCredentials authmodel.UserCredentials) (userAuthInfo authmodel.UserAuthInfo, err error)
}

type userCredentialsStorage struct {
	store  UserCredentialsStorage
	uStore userrepo.FindByUsernameStorage
}

func NewUserCredentials(store UserCredentialsStorage, uStore userrepo.FindByUsernameStorage) *userCredentialsStorage {
	return &userCredentialsStorage{
		store:  store,
		uStore: uStore,
	}
}

func (ucs *userCredentialsStorage) Login(ctx context.Context, userCredentials authmodel.UserCredentials) (authmodel.UserAuthInfo, error) {
	var user = new(usermodel.User)
	var authInfo = authmodel.NewUserAuthInfo()
	var err error

	user, err = ucs.uStore.FindUserByUsername(ctx, userCredentials.Username)

	if err != nil {
		return authInfo, err
	}

	if ok := utils.CheckPassword(userCredentials.Password, user.Password); !ok {
		return authInfo, common.ErrWithMessage(err, "Username or Password is invalid")
	}

	tokenProvider := jwt.NewTokenProvider(token.WithPathToPrivateKey("/keys/priv"))
	authInfo.AccessToken, _ = tokenProvider.Generate(*user, token.WithExpiry(time.Minute*15))
	authInfo.RefreshToken, _ = tokenProvider.Generate(*user, token.WithExpiry(time.Hour*24*7))
	authInfo.User = user

	return authInfo, nil
}
