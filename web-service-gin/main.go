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
	{ID: "1", Title: "Dragostea Din Tei", Artist: "Song by O-Zone", Price: 6.7},
	{ID: "2", Title: "Hear Me Tonight", Artist: "Modjo", Price: 1.99},
	{ID: "3", Title: "Rip & Tear", Artist: "Mick Gordon", Price: 9.99},
	{ID: "4", Title: "The Pretender", Artist: "Foo Fighters", Price: 20},
}

func main() {
	router := gin.Default()

	router.GET("/songs", getSongs)

	router.Run("localhost:8080")
}

func getSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}
