package secumodel

type CreateSecuUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (CreateSecuUser) TableName() string {
	return "users"
}