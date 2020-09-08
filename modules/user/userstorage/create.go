package userstorage

import (
	"context"
	common "go-rest-api/common/errors"
	"go-rest-api/modules/user/usermodel"
)

func (store *userSQLStorage) CreateUser(ctx context.Context, createUser usermodel.User) (usermodel.SimpleUser, error) {
	db := store.SQL.New().Begin()

	if err := db.Table(usermodel.User{}.TableName()).Create(&createUser).Error; err != nil {
		db.Rollback()
		return usermodel.SimpleUser{}, common.ErrDb(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return usermodel.SimpleUser{}, common.ErrDb(err)
	}
	return createUser.ToSimpleUser(), nil
}
