package userstorage

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"go-rest-api/modules/user/usermodel"
)

func (store *userSQLStorage) FindUserByUsername1(ctx context.Context, user usermodel.User) (usermodel.User, error) {
	db := store.SQL.New().Begin()

	var u usermodel.User
	if err := db.Table(usermodel.User{}.TableName()).Find(&u).
		Where("username = ? and password = ?", user.Username, user.Password).
		Limit(1).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//TODO: need to be handle not found here

		}
		return u, errors.New("cannot find user")
	}

	return u, nil
}
