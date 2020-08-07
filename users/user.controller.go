package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-rest-api/db"
	"go-rest-api/jwt"
	"go-rest-api/types"
	"go-rest-api/utils"
)

func Register(c *gin.Context) {

	var u User
	if err := c.Bind(&u); err != nil {
		c.Header("Content-Type", "application/json")
		utils.HttpResponBadRequest(c, types.JsonResponse{
			"message": err.Error(),
		})
		return
	}

	u.Password, _ = utils.HashPassword(u.Password)

	conn := db.GetConnection()
	collection := conn.Database("2land").Collection("users")
	result, err := collection.InsertOne(context.TODO(), u)

	if err != nil {
		panic(err.Error())
	}

	utils.HttpResponseOk(c, types.JsonResponse{
		"message": "register success",
		"userId": result.InsertedID,
	})
	return
}

func Login(c *gin.Context) {

	var userFromClient User
	if err := c.Bind(&userFromClient); err != nil {
		c.Header("Content-Type", "application/json")
		utils.HttpResponBadRequest(c, types.JsonResponse{
			"message": err.Error(),
		})
		return
	}

	userFromDb := findByUsername(userFromClient.Username)
	if isMatch := utils.CheckPassword(userFromClient.Password, userFromDb.Password); !isMatch {
		utils.HttpResponBadRequest(c, types.JsonResponse{
			"message": "username or password is incorrect",
		})
		return
	}

	token, _ := jwt.SignToken(userFromDb.Id)
	utils.HttpResponseOk(c, types.JsonResponse {
		"token": token,
	})
	return
}
