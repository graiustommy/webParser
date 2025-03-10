package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"webParser/internal/storage"
)

func StartAPIServer(db *sqlx.DB, host string) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		products := storage.GetAllProducts(db)
		c.JSON(http.StatusOK, gin.H{"products": products})
	})
	err := r.Run(host)
	if err != nil {
		log.Fatal(err)
	}
}
