package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
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

	// Generate the QR code as a PNG image in memory
	qrCode, err := qrcode.Encode(ipAddress, qrcode.Medium, 256) // 256x256 pixels
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Send the QR code image as a response
	ctx.Data(http.StatusOK, "image/png", qrCode)
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
