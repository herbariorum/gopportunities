package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	// Define your API routes here
	{
		api.GET("/opportunities", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List of opportunities",
			})
		})

		api.POST("/opportunities", func(c *gin.Context) {
			c.JSON(201, gin.H{
				"message": "Opportunity created",
			})
		})

		api.GET("/opportunities/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{
				"message": "Details of opportunity " + id,
			})
		})

		api.PUT("/opportunities/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{
				"message": "Opportunity " + id + " updated",
			})
		})

		api.DELETE("/opportunities/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{
				"message": "Opportunity " + id + " deleted",
			})
		})
	}
}
