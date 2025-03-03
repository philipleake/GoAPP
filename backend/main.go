package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Movie struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"name"`
	Director string `json:"director"`
}

var db *gorm.DB

func initDB() {
	var err error
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	info := "host=localhost user=postgres password=" + os.Getenv("PWD") + " dbname=movies port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(info), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	db.AutoMigrate(&Movie{})
}

func main() {
	initDB()
	r := gin.Default()

	r.GET("/movies", getMovies)
	r.POST("/movies", addMovie)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getMovies(c *gin.Context) {
	var movies []Movie
	db.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func addMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&movie)
	c.JSON(http.StatusCreated, movie)
}
