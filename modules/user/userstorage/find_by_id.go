package userstorage

import (
	"context"
	"errors"
	"go-rest-api/modules/user/usermodel"
)

func (store *userSQLStorage) FindUserById(ctx context.Context, userId string) (usermodel.User, error) {
	var user usermodel.User
	db := store.SQL.New().Begin()

	if err := db.Table(usermodel.User{}.TableName()).Find(&user).Where("id = ?", userId).Limit(1).Error; err != nil {
		return usermodel.User{}, errors.New("cannot find user")
	}

	return user, nil
}
