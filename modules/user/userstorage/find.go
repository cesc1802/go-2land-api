package userstorage

import (
	"context"
	"errors"
	"go-rest-api/modules/user/usermodel"
)

func (u *userSQLStorage) FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error) {

	var user usermodel.User
	db := u.SQL.New().Begin()

	if err := db.Table(usermodel.User{}.TableName()).Find(&user).Where("username = ?", userName).Limit(1).Error; err != nil {
		return nil, errors.New("cannot find user")
	}

	return &user, nil
}
