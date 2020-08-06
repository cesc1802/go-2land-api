package users

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api/db"
	"go-rest-api/jwt"
	"go-rest-api/utils"
	"net/http"
)

func Register(c *gin.Context) {

	var u User
	if err := c.Bind(&u); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	u.Password, _ = utils.HashPassword(u.Password)

	conn := db.GetConnection()
	collection := conn.Database("2land").Collection("users")
	result, err := collection.InsertOne(context.TODO(), u)

	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
		"data":    result.InsertedID,
	})
}

func Login(c *gin.Context) {

	var userFromClient User
	if err := c.Bind(&userFromClient); err != nil {
		c.Header("Content-Type", "application/json")
		utils.HttpResponBadRequest(c, err.Error())
		return
	}

	userFromDb := findByUsername(userFromClient.Username)
	if isMatch := utils.CheckPassword(userFromClient.Password, userFromDb.Password); !isMatch {
		utils.HttpResponBadRequest(c, "username or password is invalid")
		return
	}

	token, _ := jwt.SignToken(userFromDb.Id)
	fmt.Println(token)
	utils.HttpResponseOk(c, struct{ token string }{token: token})
	return
}
