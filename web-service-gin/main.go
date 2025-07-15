package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)



	router.Run("localhost:8080")
	fmt.Println("Server is running on : http://localhost:8080s")
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(cxt *gin.Context) {
	cxt.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(cxt *gin.Context){
	var newAlbum album

	if err := cxt.BindJSON(&newAlbum); err != nil {
		cxt.IndentedJSON(http.StatusBadRequest, newAlbum)
		return 
	}

	albums = append(albums, newAlbum)
	cxt.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
func getAlbumByID(cxt *gin.Context){
	id := cxt.Param("id")
	fmt.Print(id)
	for _, album := range albums {
		if album.ID == id {
			cxt.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	cxt.IndentedJSON(http.StatusNotFound, gin.H{"message" : "album not found "})
}







