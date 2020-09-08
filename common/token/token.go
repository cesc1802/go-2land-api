package token

import (
	"errors"
	"go-rest-api/modules/user/usermodel"
	"time"
)

var (
	// ErrNotFound is returned when a token cannot be found
	ErrNotFound = errors.New("token not found")
	// ErrEncodingToken is returned when the service encounters an error during encoding
	ErrEncodingToken = errors.New("error encoding the token")
	// ErrInvalidToken is returned when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
)

type JwtPayload struct {
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}

// Provider generates and inspects tokens
type Provider interface {
	//Generate(userId, roleId string, opts ...GenerateOption) (*Token, error)
	Generate(user usermodel.User, opts ...GenerateOption) (*Token, error)
	Inspect(token string) (*JwtPayload, error)
	String() string
}

type Token struct {
	// The actual token
	Token string `json:"token"`
	// Time of token creation
	Created time.Time `json:"created"`
	// Time of token expiry
	Expiry time.Time `json:"expiry"`
}
