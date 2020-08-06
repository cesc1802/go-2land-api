package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"go-rest-api/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

func SignToken(userId primitive.ObjectID) (string, error) {
	var token string
	var err error
	var expTime int64
	expTime , _ = strconv.ParseInt(config.JwtConf.JwtExpired, 10, 0)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id": userId,
		"exp": time.Now().Add(time.Minute * time.Duration(expTime)).Unix(),
	})

	token, err = at.SignedString([]byte(config.JwtConf.JwtSecret))

	if err != nil {
		return "", err
	}

	return token, err
}