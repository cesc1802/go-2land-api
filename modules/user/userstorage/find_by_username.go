package userstorage

import (
	"context"
	"github.com/jinzhu/gorm"
	common "go-rest-api/common/errors"
	"go-rest-api/modules/user/usermodel"
)

func (store *userSQLStorage) FindUserByUsername(ctx context.Context, userName string) (*usermodel.User, error) {
	db := store.SQL.New().Begin()
	var u usermodel.User

	if err := db.Table(usermodel.SimpleUser{}.TableName()).Where("username = ?", userName).First(&u).Error; err != nil {
		db.Rollback()
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound(err, "Username or Password is invalid")
		}
		return nil, common.ErrDb(err)
	}
	return &u, nil
}
