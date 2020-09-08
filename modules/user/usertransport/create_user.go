package usertransport

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/modules/user/userhdl"
	"go-rest-api/modules/user/usermodel"
	"go-rest-api/modules/user/userrepo"
	"go-rest-api/modules/user/userstorage"
	"net/http"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser usermodel.CreateUser
		var err error

		if err = c.ShouldBind(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//TODO: setup dependencies
		userStorage := userstorage.NewUserSQLStorage(db)
		userRepo := userrepo.NewCreateUserStorage(userStorage)
		userHdl := userhdl.NewCreateUserHdl(userRepo)

		createdUser, err := userHdl.CreateUser(c.Request.Context(), newUser)
		if err != nil {
			c.Error(err).SetType(1000)
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"user": createdUser,
		})
		return
	}
}
