package usermodel

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}


