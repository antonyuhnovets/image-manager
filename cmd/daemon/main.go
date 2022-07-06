package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/antonyuhnovets/image-manager/pkg/config"
	"github.com/antonyuhnovets/image-manager/pkg/middlewares"
	"github.com/antonyuhnovets/image-manager/pkg/routes"
)

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()
	fmt.Println(cfg.AMQP_URL)

	p := &broker.Producer{}
	c := &broker.Consumer{}

	p.Connect(cfg)
	c.Connect(cfg)
	go c.Consume()

	router.Use(middlewares.BrokerMiddleware(c, p))

	routes.GetRoutes(router)

	router.Run(":8080")
}
