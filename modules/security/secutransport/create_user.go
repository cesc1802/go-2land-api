package secutransport

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/modules/security/secuhdl"
	models "go-rest-api/modules/security/secumodel"
	"go-rest-api/modules/security/securepo"
	storage "go-rest-api/modules/security/secustorage"
	"net/http"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonUSer models.CreateSecuUser

		if err := c.Bind(&jsonUSer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		//setup dependency
		store := storage.NewSecurityStorage(db)
		repo := securepo.NewCreateUserStorage(store)
		handler := secuhdl.NewCreateUserHdl(repo)

		toUser := models.ToSecurityUser(jsonUSer)
		if err := handler.CreateUser(c.Request.Context(), toUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "create user successed",
		})
		return
	}
}
