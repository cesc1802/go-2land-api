package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func Login(c *gin.Context) {
	type User struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	var userJson User

	if err := c.Bind(&userJson); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": userJson,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userJson,
		})
	}
}

func MakeToken(userId uint64) (string, error) {
	var err error

	os.Setenv("JWT_SECRET", "abcdefgh1234")
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id": userId,
		"exp_time": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return token, err
}

