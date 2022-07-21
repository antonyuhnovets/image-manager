package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/storage"
)

// Handler for downloading image from storage
// Used by download route to process request,
// search for file in storage and return responce
func DownloadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		storage, ok := c.MustGet("Storage").(storage.Entity)
		if !ok {
			log.Fatalf("Context doesn`t have consumer")
		}

		if c.Param("img_id") == "" {
			c.JSON(400, "Bad request")
		}

		switch c.Param("quality") {
		case "25", "50", "75", "100":
			img := fmt.Sprintf(
				"%s_%s",
				c.Param("quality"),
				c.Param("img_id"),
			)
			fpath, err := storage.GetImgByName(img)
			if err != nil {
				log.Println(err)
				c.JSON(400, "Bad request")
			}
			c.File(fpath)

		default:
			c.JSON(400, "Bad request")
		}
	}
}
