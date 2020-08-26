package authtransport

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/modules/auth/authmodel"
	"go-rest-api/modules/auth/authrepo"
	"go-rest-api/modules/auth/authstorage"
	"go-rest-api/modules/user/userstorage"
	"net/http"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonUSer authmodel.LoginUser

		if err := c.Bind(&jsonUSer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		//TODO: setup dependency here
		authStore := authstorage.NewAuthSQLStorage(db)
		userStore := userstorage.NewUserSQLStorage(db)
		authRepo := authrepo.NewLoginUserStorage(userStore, authStore)

		return

	}
}
