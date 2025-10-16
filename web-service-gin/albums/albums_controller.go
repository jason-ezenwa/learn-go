package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAlbumController() *AlbumController {
	return &AlbumController{service: NewAlbumService()}
}

func (controller *AlbumController) PostAlbum(context *gin.Context) {
	var newAlbum Album

	err := context.BindJSON(&newAlbum)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdAlbum := controller.service.PostAlbum(newAlbum)

	context.IndentedJSON(http.StatusCreated, createdAlbum)
}

func (controller *AlbumController) GetAlbumByID(context *gin.Context) {
	id := context.Param("id")

	album := controller.service.GetAlbumByID(id)

	if album == nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, album)
}

func (controller *AlbumController) GetAlbums(context *gin.Context) {
	albums := controller.service.GetAlbums()

	context.IndentedJSON(http.StatusOK, albums)
}

func (controller *AlbumController) RegisterRoutes(router *gin.Engine) {
	// Register routes for the albums group
	albums := router.Group("/albums")
	{
		albums.GET("", controller.GetAlbums)
		albums.POST("", controller.PostAlbum)
		albums.GET("/:id", controller.GetAlbumByID)
	}
}