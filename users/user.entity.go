package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(50);not null"`
}

func (User) TableName() string {
	return "users"
}