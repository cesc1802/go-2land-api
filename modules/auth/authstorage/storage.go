package authstorage

import (
	"github.com/jinzhu/gorm"
)

type authSQLStorage struct {
	SQL *gorm.DB
}

func NewAuthSQLStorage(SQL *gorm.DB) *authSQLStorage {
	return &authSQLStorage{
		SQL: SQL,
	}
}