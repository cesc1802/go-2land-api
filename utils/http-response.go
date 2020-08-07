package utils

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/types"
	"net/http"
)

func HttpResponseOk( c *gin.Context, response types.JsonResponse) {
	c.JSON(http.StatusOK, gin.H {
		"data": response,
	})
}

func HttpResponBadRequest(c *gin.Context, response types.JsonResponse) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data": response,
	})
}

func HttpResponseInternalServerError(c *gin.Context, response types.JsonResponse) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"data": response,
	})
}