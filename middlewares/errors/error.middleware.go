package errormiddleware

import (
	"github.com/gin-gonic/gin"
	common "go-rest-api/common/errors"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		//skip if no errors
		//if c.Errors.Last() == nil {
		//	return
		//}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case 404:
					err := e.Err.(common.AppError)
					//c.AbortWithError(err.ErrorCode, err)
					c.JSON(err.ErrorCode, gin.H{
						"error":   err.Error(),
						"message": "Not Found",
					})
					return
				case 400:
					err := e.Err.(common.AppError)
					//c.AbortWithError(err.ErrorCode, err)
					c.JSON(err.ErrorCode, gin.H{
						"error":   err.Error(),
						"message": "Bad Request",
					})
					return
				case 500:
					err := e.Err.(common.AppError)
					//c.AbortWithError(err.ErrorCode, err)
					c.JSON(err.ErrorCode, gin.H{
						"error":   err.Error(),
						"message": "Server internal error",
					})
					return
				}

			}
		}
	}
}
