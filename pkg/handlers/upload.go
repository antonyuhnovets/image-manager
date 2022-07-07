package handlers

import (
	"log"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/gin-gonic/gin"
)

func UploadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod, ok := c.MustGet("Producer").(*broker.Producer)
		if !ok {
			log.Fatalf("Context doesn`t have producer")
		}
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, "Bad request")
		}
		if file == nil {
			c.JSON(400, "Bad request")
		}
		format := header.Header.Get("Content-Type")
		if format == "image/jpeg" || format == "image/png" {
			c.JSON(200, "File accepted")
		} else {
			c.JSON(400, "Bad request: unknown type")
		}
		err = prod.Publish(file, header)
		if err != nil {
			c.JSON(400, err)
		}
	}
}
