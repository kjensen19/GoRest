package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// json: is what it will be serialized into, otherwise will default to capitalized field names-- not common in JSON
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//$ go run . to start server

//GET request in terminal:
// $ curl http://localhost:8080/albums \
//     --header "Content-Type: application/json" \
//     --request "GET"

//GET single album by ID
// curl http://localhost:8080/albums/2 (id is 2)

//POST request example:
// $ curl http://localhost:8080/albums \
// --include \
// --header "Content-Type: application/json" \
// --request "POST" \
// --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GET
// context carries all of the important stuff (request details, validation and serialization of JSON, etc)
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//Returns 200(OK)

//POST

func postAlbums(c *gin.Context) {
	var newAlbum album
	//Calling bind JSON attaches the received to the newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//Add newAlbum to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	//Returns 201 created
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	//retrieve id from request
	id := c.Param("id")

	//Loop through albums looking for matching id, if found serialize and return as JSON along with 200 OK
	//Usually this would be a DB query
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//returns 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
