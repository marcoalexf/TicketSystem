package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/qrcode", func(ctx *gin.Context) {
		ipAddress := ctx.ClientIP()

		ctx.JSON(http.StatusOK, gin.H{
			"ipAddress": ipAddress,
		})
	})

	router.GET("/ticket", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ticket": 1,
		})
	})

	// Start the server
	router.Run(":8080") // Listen on port 8080
}
