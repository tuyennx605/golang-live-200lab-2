package ginitem

import (
	"net/http"
	"strconv"
	"todo-list/common"
	"todo-list/module/item/biz"
	"todo-list/module/item/model"
	"todo-list/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester) // lấy curent user. đưa về requester

		store := storage.NewSQLStore(db)
		bussiness := biz.NewUpdateItemBiz(store, requester)

		if err := bussiness.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
