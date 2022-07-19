// Start server with all components

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/antonyuhnovets/image-manager/pkg/compressor"
	"github.com/antonyuhnovets/image-manager/pkg/config"
	"github.com/antonyuhnovets/image-manager/pkg/middlewares"
	"github.com/antonyuhnovets/image-manager/pkg/routes"
	"github.com/antonyuhnovets/image-manager/pkg/storage"
)

func main() {
	// Set router
	router := gin.Default()

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Set storage and compressor tool for imgs
	s := storage.GetStorage(&cfg)
	c := compressor.GetCompressor(s.SaveImg)

	// Connect to message broker, start consuming
	b := broker.SetBroker(&cfg)
	defer b.Producer.Disconnect()
	defer b.Consumer.Disconnect()
	go b.Consumer.Consume(c.Handler)

	// Use broker and storage middleware
	router.Use(middlewares.Middleware(b, s))

	// Get API endpoints
	routes.GetRoutes(router)

	// Start server
	router.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
