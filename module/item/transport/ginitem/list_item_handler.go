package ginitem

import (
	"net/http"
	"todo-list/common"
	"todo-list/module/item/biz"
	"todo-list/module/item/model"
	"todo-list/module/item/repository"
	"todo-list/module/item/storage"
	"todo-list/module/item/storage/rpc"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()
		requester := c.MustGet(common.CurrentUser).(common.Requester) // lấy curent user. đưa về requester

		// old : internal call
		// store := storage.NewSQLStore(db)
		// likeStore := userlikestore.NewSQLStore(db)
		// repo := repository.NewListItemRepo(store, likeStore, requester)
		// business := biz.NewListItemRepo(repo, requester)

		// // new: call multiple service call = resful api
		// store := storage.NewSQLStore(db)
		// likeStore := resapi.New("http://localhost:3005")
		// repo := repository.NewListItemRepo(store, likeStore, requester)
		// business := biz.NewListItemRepo(repo, requester)

		// new: call multiple service call = GRPC
		store := storage.NewSQLStore(db)
		likeStore := rpc.NewClient()
		repo := repository.NewListItemRepo(store, likeStore, requester)
		business := biz.NewListItemRepo(repo, requester)

		result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
