package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/antonyuhnovets/image-manager/pkg/handlers"
)

// Set up API endpoints
func GetRoutes(router *gin.Engine) {
	router.POST("/upload/:file", handlers.UploadImg())
	router.GET("/download/:img_id/:quality", handlers.DownloadImg())
}
