package usermodel

import (
	"github.com/google/uuid"
	"go-rest-api/utils"
)

type CreateUser struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (CreateUser) TableName() string {
	return "users"
}

func (u CreateUser) ToUser() User {
	hashPass, _ := utils.HashPassword(u.Password)
	return User{
		ID:       uuid.New().String(),
		Username: u.Username,
		Password: hashPass,
	}
}
