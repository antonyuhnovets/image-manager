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
	// Set router
	router := gin.Default()

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Declare producer and consumer
	p := &broker.Producer{}
	c := &broker.Consumer{}

	// Connect message broker
	p.Connect(cfg)
	defer p.Disconnect()
	c.Connect(cfg)
	defer c.Disconnect()
	go c.Consume()

	// Use broker middleware
	router.Use(middlewares.BrokerMiddleware(c, p))

	// Get API endpoints
	routes.GetRoutes(router)

	// Start server
	router.Run(":8080")
}
