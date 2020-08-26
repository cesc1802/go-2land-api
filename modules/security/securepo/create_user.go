package securepo

import (
	"context"
	"errors"
	models "go-rest-api/modules/security/secumodel"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserStorage interface {
	Create(ctx context.Context, user models.Security) error
}

type createUserStorage struct {
	userStorage CreateUserStorage
}

func NewCreateUserStorage(userStorage CreateUserStorage) *createUserStorage {
	return &createUserStorage{
		userStorage: userStorage,
	}
}

func (c *createUserStorage) Create(ctx context.Context, user models.Security) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return errors.New(err.Error())
	}

	user.Password = string(hashPassword)
	return c.userStorage.Create(ctx, user)
}
