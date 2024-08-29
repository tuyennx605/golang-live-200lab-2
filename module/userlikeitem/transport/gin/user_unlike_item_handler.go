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

func UnLikeItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester) // lấy curent user. đưa về requester

		store := storage.NewSQLStore(db)
		business := biz.NewUserUnLikeItemBiz(store)

		if err := business.UnLikeItem(c.Request.Context(), requester.GetUserId(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
