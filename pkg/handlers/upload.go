package handlers

import (
	"fmt"
	"log"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/gin-gonic/gin"
)

func UploadImg() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Producer\n", c.Value("Producer"))
		prod, ok := c.MustGet("Producer").(*broker.Producer)
		if !ok {
			log.Fatalf("Context doesn`t have producer")
		}
		fmt.Println("Producer accepted")
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, "Bad request")
		}
		fmt.Printf("File accepted %s\n", file)
		switch file {
		case nil:
			c.JSON(400, "Bad request")
		default:
			prod.Publish(file, header)
			fmt.Println("Handler complete")
		}
	}
}
