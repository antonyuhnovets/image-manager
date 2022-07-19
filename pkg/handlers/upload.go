package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
)

// Handler for uploading image route
// Takes request with file as value for "file" param
// Generates uniq id and call broker's publish method
func UploadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		broker, ok := c.MustGet("Broker").(*broker.Broker)
		if !ok {
			log.Fatalf("Context doesn`t have producer")
		}

		file, header, err := c.Request.FormFile("file")
		if err != nil || file == nil {
			c.JSON(400, "Bad request")
		}

		id := strconv.Itoa(int(uuid.New().ID()))
		format := header.Header.Get("Content-Type")

		if format == "image/jpeg" || format == "image/png" {
			c.JSON(200, fmt.Sprintf("File with id %s accepted", id))
		} else {
			c.JSON(400, "Bad request: unknown type")
		}

		err = broker.Producer.Publish(id, file, header)
		if err != nil {
			c.JSON(400, err)
		}
	}
}
