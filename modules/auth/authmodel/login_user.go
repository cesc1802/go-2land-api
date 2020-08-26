package authmodel

type LoginUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func(LoginUser) TableName() string {
	return "users"
}