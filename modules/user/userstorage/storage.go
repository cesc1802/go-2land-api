package userstorage

import (
	"github.com/jinzhu/gorm"
)

type userSQLStorage struct {
	SQL *gorm.DB
}

func  NewUserSQLStorage(db *gorm.DB) *userSQLStorage {
	return &userSQLStorage{
		SQL: db,
	}
}
