package usermodel

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"role_id"`
}

func (User) TableName() string {
	return "users"
}

func (u User) ToSimpleUser() SimpleUser {
	return SimpleUser{
		ID:       u.ID,
		Username: u.Username,
		RoleId:   u.RoleId,
	}
}