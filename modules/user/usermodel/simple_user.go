package usermodel


type SimpleUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RoleId   string `json:"role_id"`
}

func (SimpleUser) TableName() string {
	return "users"
}