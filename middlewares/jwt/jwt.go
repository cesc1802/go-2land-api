package jwtmiddleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

type JWTMiddleware interface {
	Inspect(ctx context.Context) gin.HandlerFunc
}

type jwtMiddleware struct {
	jwt       JWTMiddleware
	TokenType string
}

func New() *jwtMiddleware {
	return &jwtMiddleware{
		TokenType: "Bearer ",
	}
}

func (j *jwtMiddleware) jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == j.TokenType) {
		return "", errors.New("invalid auth header")
	}
	return parts[1], nil
}

func (j *jwtMiddleware) jwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)

	if cookie == "" {
		return "", errors.New("empty cookie")
	}
	return cookie, nil
}

func (j *jwtMiddleware) Inspect(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := j.jwtFromHeader(c, "Authorization")
		c.Next()
	}
	return j.jwt.Inspect(ctx)
}

func (jmw *jwtMiddleware) ParseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(jwt.SigningMethodHS256) != t.Method {
			return nil, errors.New("invalid signature algorithm")
		}
	})
}

func (j *jwtMiddleware) Generate(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
