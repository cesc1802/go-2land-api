package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/users"
	"log"
	"net/http"
)

func homeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome home!")
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Static("/public", "./public")
	client := r.Group("/api/2land")
	{
		client.POST("/auth/login", users.Login)
		client.POST("/auth/register", users.Register)
	}
	return r
}

func main() {
	err:=godotenv.Load()

	if err != nil {
		log.Fatalln("error loading .env file")
	}
	config.Read()
	config.ReadDbConf()
	db.Init()



	r := setupRouter()
	r.Run(":3000")
}
