package authtransport

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	common "go-rest-api/common/errors"
	"go-rest-api/modules/auth/authhdl"
	"go-rest-api/modules/auth/authmodel"
	"go-rest-api/modules/auth/authrepo"
	"go-rest-api/modules/auth/authstorage"
	"go-rest-api/modules/user/userstorage"
	"net/http"
)

func DoLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userCredentials authmodel.UserCredentials
		var err error
		if err = c.Bind(&userCredentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		authStore := authstorage.NewAuthSQLStorage(db)
		userStore := userstorage.NewUserSQLStorage(db)
		userCredentialsRepo := authrepo.NewUserCredentials(authStore, userStore)
		userCredentialsHandler := authhdl.NewUserCredentialsHdl(userCredentialsRepo)

		var authUserInfo authmodel.UserAuthInfo
		authUserInfo, err = userCredentialsHandler.Login(c.Request.Context(), userCredentials)

		if err != nil {
			appError := err.(common.AppError)
			c.Error(err).SetType(appError.ErrType)
			return
		}

		c.SetCookie("rotk", authUserInfo.RefreshToken.Token, 3600, "", "", false, false)
		c.JSON(http.StatusOK, gin.H{
			"access_token": authUserInfo.AccessToken.Token,
			"user_id":      authUserInfo.User.ID,
		})
		return
	}
}
