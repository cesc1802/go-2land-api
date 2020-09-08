package jwt

import (
	"fmt"
	"go-rest-api/modules/user/usermodel"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go-rest-api/common/token"
)

// authClaims to be encoded in the JWT
type authClaims struct {
	//UserId string `json:"user_id"`
	//RoleId string `json:"role_id"`

	Payload token.JwtPayload `json:"payload"`
	jwt.StandardClaims
}

// JWT implementation of token provider
type JWT struct {
	opts token.Options
}

// NewTokenProvider returns an initialized basic provider
func NewTokenProvider(opts ...token.Option) token.Provider {
	return &JWT{
		opts: token.NewOptions(opts...),
	}
}

// Generate a new JWT
func (j *JWT) Generate(user usermodel.User, opts ...token.GenerateOption) (*token.Token, error) {
	// decode the private key
	pwd, _ := os.Getwd()
	if j.opts.PathToPrivateKey == "" {
		j.opts.PathToPrivateKey = "/keys/priv"
	}
	keyPath := pwd + j.opts.PathToPrivateKey
	priv, err := ioutil.ReadFile(keyPath)

	// parse the private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, token.ErrEncodingToken
	}

	// parse the options
	options := token.NewGenerateOptions(opts...)

	// generate the JWT
	expiry := time.Now().Add(options.Expiry)

	fmt.Println(expiry.String())
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, authClaims{
		token.JwtPayload{
			UserId: user.ID,
			RoleId: user.RoleId,
		},
		jwt.StandardClaims{
			Subject:   user.ID,
			ExpiresAt: expiry.Unix(),
		},
	})
	tok, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}

	// return the token
	return &token.Token{
		Token:   tok,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

// Inspect a JWT
func (j *JWT) Inspect(t string) (*token.JwtPayload, error) {
	// decode the public key
	if j.opts.PathToPublicKey == "" {
		j.opts.PathToPublicKey = "/keys/pub"
	}
	pwd, _ := os.Getwd()
	fullPath := pwd + j.opts.PathToPublicKey
	pub, err := ioutil.ReadFile(fullPath)

	// parse the public key
	res, err := jwt.ParseWithClaims(t, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(pub)
	})

	if err != nil {
		return nil, token.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, token.ErrInvalidToken
	}
	claims, ok := res.Claims.(*authClaims)
	if !ok {
		return nil, token.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

// String returns JWT
func (j *JWT) String() string {
	return "jwt"
}
