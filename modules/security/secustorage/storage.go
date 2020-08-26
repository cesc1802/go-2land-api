package security

import "github.com/jinzhu/gorm"

type securityStorage struct {
	db *gorm.DB
}

func NewSecurityStorage(db *gorm.DB) *securityStorage {
	return &securityStorage{
		db: db,
	}
}
