package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
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

func GetDB() *gorm.DB {
	return db
}
