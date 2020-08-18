package poststorage

import (
	"github.com/jinzhu/gorm"
)

type postSQLStorage struct {
	db *gorm.DB
}

func NewPostSQLStorage(db *gorm.DB) *postSQLStorage {
	return &postSQLStorage{db: db}
}
