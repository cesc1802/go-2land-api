package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/modules/post/posttransport"
	"go-rest-api/modules/security/secutransport"
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

	r := gin.New()
	r.Use(gin.Logger())

	client := r.Group("/api/2land")
	{
		client.GET("/posts", posttransport.GetAll(dbConn.GetConnection()))
		client.POST("/posts", posttransport.CreatePost(dbConn.GetConnection()))
		client.Group("/auth").POST("/register", secutransport.CreateUser(dbConn.GetConnection()))
	}

	r.Run(":3000")

}
