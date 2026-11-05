package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var songs = []song{
	{ID: "1", Title: "Dragostea Din Tei", Artist: "O-Zone", Price: 6.7},
	{ID: "2", Title: "Hear Me Tonight", Artist: "Modjo", Price: 1.99},
	{ID: "3", Title: "Rip & Tear", Artist: "Mick Gordon", Price: 9.99},
	{ID: "4", Title: "The Pretender", Artist: "Foo Fighters", Price: 20},
}

func main() {
	router := gin.Default()

	router.GET("/songs", getSongs)
	router.GET("/songs/:id", getSongByID)
	router.POST("/songs", postSong)

	router.Run("localhost:8080")
}

func getSongs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, songs)
}

func postSong(context *gin.Context) {
	var newSong song

	// Call BindJSON to bind the received JSON to newSong
	if err := context.BindJSON(&newSong); err != nil {
		return
	}

	// Add the new song to the slice
	songs = append(songs, newSong)
	context.IndentedJSON(http.StatusCreated, newSong)
}

func getSongByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range songs {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
}
