package upload

import (
	"fmt"
	"net/http"
	"time"
	"todo-list/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfill("http://localhost:3005")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
