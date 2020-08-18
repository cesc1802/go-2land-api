package posttransport

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/modules/post/posthandler"
	"go-rest-api/modules/post/postmodel"
	repo2 "go-rest-api/modules/post/postrepo"
	storage "go-rest-api/modules/post/poststorage"
	"net/http"
)

func GetAll(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []postmodel.Post
		var err error

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		store := storage.NewPostSQLStorage(db)
		repo := repo2.NewGetPostStorage(store)
		hdl := posthandler.NewGetPostHdl(repo)

		if data, err = hdl.GetAll(c.Request.Context()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    data,
		})
		return
	}
}
