package usergin

import (
	"net/http"
	"todo-list/common"
	"todo-list/module/user/biz"
	"todo-list/module/user/model"
	"todo-list/module/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()
		biz := biz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}
