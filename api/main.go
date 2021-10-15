package main

import (
	"log"
	"net/http"

	"api/main.go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server
	router.Run(":3000")

	db, err := gorm.Open(sqlite.Open("adaDb1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.New()
	r.
		GET("/promos", handler.ListPromoHandler(db))
	srv := &http.Server{
		Handler: r,
	}

	go func() {
		log.Printf("Http server is started on interface %s:%d")

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
