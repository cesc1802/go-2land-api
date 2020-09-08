package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rest-api/config"
	"go-rest-api/db"
	errormiddleware "go-rest-api/middlewares/errors"
	jwtmiddleware "go-rest-api/middlewares/jwt"
	"go-rest-api/modules/auth/authtransport"
	"go-rest-api/modules/post/posttransport"
	"go-rest-api/modules/user/usertransport"
	"log"
	_ "net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	config.Read()
	config.ReadDbConf()
	dbConn := db.New()

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(errormiddleware.ErrorHandle())

	jwtVerifier := jwtmiddleware.NewJwtMiddleware(dbConn.GetConnection())

	client := app.Group("/api/2land")

	postGroup := client.Group("").Use(jwtVerifier.Verify())
	{
		postGroup.GET("/posts", jwtVerifier.Verify(), posttransport.GetAll(dbConn.GetConnection()))
		postGroup.POST("/posts", posttransport.CreatePost(dbConn.GetConnection()))

	}

	authGroup := client.Group("/auth")
	{
		authGroup.POST("/login", authtransport.DoLogin(dbConn.GetConnection()))
	}

	userGroup := client.Group("/users")
	{
		userGroup.POST("/register", usertransport.CreateUser(dbConn.GetConnection()))
	}

	app.Run(":3001")

}

type Widget interface {
	ID() string
}

type widget struct {
	id string
}

func NewWidget() Widget{
	return widget{
		id: "1234",
	}
}

func (w widget) ID() string {
	return w.id
}
