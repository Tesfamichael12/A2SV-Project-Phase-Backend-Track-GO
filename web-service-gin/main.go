package main

import (
	"fmt"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example/web-service-gin/docs"

	"github.com/gin-gonic/gin"
)

// album represents a music album.
// @Description Album object
type album struct {
	ID     string  `json:"id" example:"1"`
	Title  string  `json:"title" example:"Blue Train"`
	Artist string  `json:"artist" example:"John Coltrane"`
	Price  float64 `json:"price" example:"56.99"`
}


var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Example format of annotation : used only on api handler functions
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string  "ok"
// @Router /string/{some_id} [get]
func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)


	// swagger API doc 
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
	fmt.Println("Server is running on : http://localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
// @Summary List albums
// @Description Get all albums in the store
// @Tags albums
// @Produce json
// @Success 200 {array} album "List of albums"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums [get]
func getAlbums(cxt *gin.Context) {
	cxt.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
// @Summary Add a new album
// @Description Add a new album to the store
// @Tags albums
// @Accept json
// @Produce json
// @Param album body album true "Album to add"
// @Success 201 {object} album "Created album"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums [post]
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
// @Summary Get album by ID
// @Description Get details of an album by its ID
// @Tags albums
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} album "Album details"
// @Failure 404 {object} map[string]string "Album not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums/{id} [get]
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







