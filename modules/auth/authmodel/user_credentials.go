package authmodel

type UserCredentials struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}