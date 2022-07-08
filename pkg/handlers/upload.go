package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/antonyuhnovets/image-manager/pkg/utils"
)

// Handler for uploading image endpoint
func UploadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod, ok := c.MustGet("Producer").(*broker.Producer)
		if !ok {
			log.Fatalf("Context doesn`t have producer")
		}

		file, header, err := c.Request.FormFile("file")
		if err != nil || file == nil {
			c.JSON(400, "Bad request")
		}

		id := utils.IdGen()
		format := header.Header.Get("Content-Type")
		if format == "image/jpeg" || format == "image/png" {
			c.JSON(200, fmt.Sprintf("File with id %s accepted", id))
		} else {
			c.JSON(400, "Bad request: unknown type")
		}

		err = prod.Publish(id, file, header)
		if err != nil {
			c.JSON(400, err)
		}
	}
}
