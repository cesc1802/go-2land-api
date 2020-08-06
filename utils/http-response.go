package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	data interface{}
}

func HttpResponseOk( c *gin.Context, data interface{}) {

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H {
		"data": data,
	})
}

func HttpResponBadRequest(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data": data,
	})
}

func HttpResponseInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"data": "",
	})
}