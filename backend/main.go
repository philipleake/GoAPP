package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"name"`
	Director string `json:"director"`
}

func main() {
	InitDB()
	r := gin.Default()

	r.GET("/movies", getMovies)
	r.POST("/movies", addMovie)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getMovies(c *gin.Context) {
	var movies []Movie
	GetDB().Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func addMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	GetDB().Create(&movie)
	c.JSON(http.StatusCreated, movie)
}
