package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"webParser/internal/models"
)

type Product struct {
	id             uint `db:"id"`
	models.Product `db:"product"`
}

func NewDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	return db
}

func SaveProducts(db *sqlx.DB, products []models.Product) {
	for _, product := range products {
		if _, err := db.Exec("INSERT INTO products (name, price) values ($1, $2)", product.Name, product.Price); err != nil {
			log.Fatalf("failed to save product: %v", err)
		}
		log.Printf("saved product: %v", product.Name)
	}
}

func GetAllProducts(db *sqlx.DB) []models.Product {
	var products []models.Product
	err := db.Select(&products, "SELECT * FROM products")
	if err != nil {
		log.Fatalf("failed to get products: %v", err)
		return nil
	}
	log.Printf("sended %d products", len(products))
	return products
}
