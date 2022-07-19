package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/broker"
	"github.com/antonyuhnovets/image-manager/pkg/storage"
)

// Middleware for passing message broker and storage
// to handler functions by putting it ot context, using by router
func Middleware(b *broker.Broker, s storage.Entity) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Broker", b)
		c.Set("Storage", s)
		c.Next()
	}
}
