package router

import (
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
	router := gin.Default()
	godotenv.Load()

	// Define a route for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
