package middlewares

import (
	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/gin-gonic/gin"
)

// Middleware for using message broker
func BrokerMiddleware(cons *broker.Consumer, prod *broker.Producer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Consumer", cons)
		c.Set("Producer", prod)
		c.Next()
	}
}
