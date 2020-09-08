package jwtmiddleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/common/token"
	"go-rest-api/common/token/jwt"
	"go-rest-api/middlewares/jwt/jwtrepo"
	"go-rest-api/modules/user/userrepo"
	"go-rest-api/modules/user/userstorage"
	"net/http"
	"strings"
)

type jwtMiddleware struct {
	verifier  jwtrepo.Verifier
	TokenType string
}

func NewJwtMiddleware(db *gorm.DB) *jwtMiddleware {

	uStorage := userstorage.NewUserSQLStorage(db)
	uRepository := userrepo.NewFindUserById(uStorage)

	return &jwtMiddleware{
		verifier:  jwtrepo.NewVerifier(uRepository),
		TokenType: "Bearer",
	}
}

func (j *jwtMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userInfo = new(token.JwtPayload)
		var err error
		tok, _ := j.jwtFromHeader(c, "Authorization")
		tokenProvider := jwt.NewTokenProvider(token.WithPathToPublicKey("/keys/pub"))
		userInfo, err = tokenProvider.Inspect(tok)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if ok := j.verifier.VerifyUser(c.Request.Context(), userInfo.UserId); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "user is invalid",
			})
			return
		}
		c.Next()
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
