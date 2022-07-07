package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/antonyuhnovets/image-manager/pkg/config"
	"github.com/antonyuhnovets/image-manager/pkg/middlewares"
	"github.com/antonyuhnovets/image-manager/pkg/routes"
)

func main() {
	router := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	p := &broker.Producer{}
	c := &broker.Consumer{}

	p.Connect(cfg)
	defer p.Disconnect()
	c.Connect(cfg)
	defer c.Disconnect()
	go c.Consume()

	router.Use(middlewares.BrokerMiddleware(c, p))

	routes.GetRoutes(router)

	router.Run(":8080")
}
