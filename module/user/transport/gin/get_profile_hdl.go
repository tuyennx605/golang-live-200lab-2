package usergin

import (
	"net/http"
	"todo-list/common"

	"github.com/gin-gonic/gin"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		// có thể dùng c.Get() trả về 2 tham số u,bool để check xem có hây ko

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
