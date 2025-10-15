package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Calloway", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range albums {
			if a.ID == id {
					context.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(context *gin.Context) {
	var newAlbum album

	// Call BindJSON to check if the request body is valid
	err := context.BindJSON(&newAlbum);

	if err != nil {
			return
	}

	albums = append(albums, newAlbum)

	context.IndentedJSON(http.StatusCreated, newAlbum)
}