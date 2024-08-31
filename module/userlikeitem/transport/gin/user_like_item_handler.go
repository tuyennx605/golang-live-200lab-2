package ginuserlikeitem

import (
	"net/http"
	"strconv"
	"todo-list/common"
	itemstorage "todo-list/module/item/storage"
	"todo-list/module/userlikeitem/biz"
	"todo-list/module/userlikeitem/model"
	"todo-list/module/userlikeitem/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeItem(db *gorm.DB) func(*gin.Context) {
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
		itemStore := itemstorage.NewSQLStore(db)
		business := biz.NewUserLikeItemBiz(store, itemStore)

		if err := business.LikeItem(c.Request.Context(), &model.Like{
			UserId: requester.GetUserId(),
			ItemId: id,
		}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
