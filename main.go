package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getQrCode handles a GET request to generate and return a QR Code for a "Giver"
// @Summary      Returns QR Code for a Giver
// @Description  Returns QR Code for a Giver
// @Tags         qrcode
// @Produce      json
// @Success      200   {object}  gin.H  "QR Code for Queue"
// @Failure      400   {object}  gin.H  "Bad Request"
// @Router       /qrcode [get]
func getQrCode(ctx *gin.Context) {
	ipAddress := ctx.ClientIP()

	ctx.JSON(http.StatusOK, gin.H{
		"ipAddress": ipAddress,
	})
}

// getQrCode handles a GET request retrieve a ticket for a given Queue
// @Summary      Returns ticket number for the queue
// @Description  Returns ticket number for the queue
// @Tags         qrcode, ticket
// @Produce      json
// @Success      200   {object}  gin.H  "Ticket number for queue"
// @Failure      400   {object}  gin.H  "Bad Request"
// @Router       /ticket [get]
func getTicket(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": 1,
	})
}

// @title           QR Code Microservice
// @version         1.0
// @description     This microservice has the responsability of creating and validating QR Codes for a Queue

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
func main() {
	router := gin.Default()

	router.GET("/qrcode", getQrCode)

	router.GET("/ticket", getTicket)

	// Start the server
	router.Run(":8080") // Listen on port 8080
}
