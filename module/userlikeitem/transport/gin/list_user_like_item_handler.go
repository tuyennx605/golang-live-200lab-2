package ginuserlikeitem

import (
	"net/http"
	"strconv"
	"todo-list/common"
	"todo-list/module/userlikeitem/biz"
	"todo-list/module/userlikeitem/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUserLikeItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var queryString struct {
			common.Paging
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storage.NewSQLStore(db)
		business := biz.ListUserLikeItemStore(store)

		result, err := business.ListUser(c.Request.Context(), id, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, nil))
	}
}
