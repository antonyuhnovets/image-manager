package routes

import (
	"github.com/antonyuhnovets/image-manager/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) {
	router.POST("/upload/:file", handlers.UploadImg())
	router.GET("/download/:img_id/:quality", handlers.DownloadImg())
}
