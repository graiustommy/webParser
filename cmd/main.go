package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"webParser/internal/api"
	"webParser/internal/models"
	"webParser/internal/storage"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var products []models.Product = []models.Product{}
	dsn := os.Getenv("DB_DSN")
	db := storage.NewDB(dsn)

	products = append(products, models.Product{ID: 0, Name: "Book", Price: "100"})
	storage.SaveProducts(db, products)
	host := os.Getenv("host")
	api.StartAPIServer(db, host)
}
