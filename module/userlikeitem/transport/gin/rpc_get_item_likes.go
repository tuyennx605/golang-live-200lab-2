package ginuserlikeitem

import (
	"net/http"
	"todo-list/common"
	"todo-list/module/userlikeitem/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItemLikes(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		type RequestData struct {
			Ids []int `json:"ids"`
		}

		var data RequestData

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		mapResult, err := store.GetItemLikes(c.Request.Context(), data.Ids)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(mapResult))
	}
}
