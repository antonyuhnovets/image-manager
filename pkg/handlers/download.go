package handlers

import (
	"fmt"
	"log"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/gin-gonic/gin"
)

func DownloadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		cons, ok := c.MustGet("Consumer").(*broker.Consumer)
		if !ok {
			log.Fatalf("Context doesn`t have consumer")
		}
		img := fmt.Sprintf("%s_%s", c.Param("quality"), c.Param("img_id"))
		fpath, err := cons.Storage.GetImgByName(img)
		if err != nil {
			c.JSON(400, "Bad request")
		}
		c.File(fpath)
	}
}