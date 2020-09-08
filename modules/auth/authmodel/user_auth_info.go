package authmodel

import (
	"go-rest-api/common/token"
	"go-rest-api/modules/user/usermodel"
)

type UserAuthInfo struct {
	User         *usermodel.User `json:"user"`
	AccessToken  *token.Token    `json:"access_token"`
	RefreshToken *token.Token    `json:"refresh_token"`
}

func NewUserAuthInfo() UserAuthInfo {
	return UserAuthInfo{
		User:         new(usermodel.User),
		AccessToken:  new(token.Token),
		RefreshToken: new(token.Token),
	}
}
