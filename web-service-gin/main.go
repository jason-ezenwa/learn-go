package main

import (
	"example/web-service-gin/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	albumController := albums.NewAlbumController()
	albumController.RegisterRoutes(router)

	router.Run("localhost:8080")
}
