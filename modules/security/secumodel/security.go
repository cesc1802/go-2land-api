package secumodel

import "github.com/google/uuid"

type Security struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Security) TableName() string {
	return "users"
}

func ToSecurityUser(user CreateSecuUser) Security {
	var userUuid uuid.UUID
	userUuid, _ = uuid.NewRandom()

	return Security{
		ID:       userUuid.String(),
		Username: user.Username,
		Password: user.Password,
	}
}
