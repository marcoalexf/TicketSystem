package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcoalexf/QRCodeService/dbconnection"
	"github.com/skip2/go-qrcode"
)

type Distributer struct {
	Email string `json:"email"`
}

// registerQueue handles a POST request to generate and return a QR Code for a "Giver" and stores the email and ipAddress in database
// @Summary      Saves in the database the ipAddress and email and Returns QR Code for a Giver
// @Description  Returns QR Code for a Giver
// @Tags         qrcode
// @Produce      json
// @Success      200   {object}  gin.H  "QR Code for Queue"
// @Failure		 400   {object}  gin.H  "Email is required or in a bad format"
// @Router       /qrcode [post]
func registerQueue(ctx *gin.Context, connection *pgxpool.Pool) {
	ipAddress := ctx.ClientIP()

	var distributer Distributer

	// Bind the JSON to the user struct
	if err := ctx.ShouldBindJSON(&distributer); err != nil {
		// Respond with an error if binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Store in database
	query := `
		INSERT INTO public.service (email, ip_address)
		VALUES ($1, $2);
	`

	_, err := connection.Exec(ctx, query, distributer.Email, ipAddress)

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
func getTicket(ctx *gin.Context, connection *pgxpool.Pool) {
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": 1,
	})
}

// @title           QR Code Microservice
// @version         1.0
// @description     This microservice has the responsability of creating and validating QR Codes for a Queue

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      localhost:4444
// @BasePath  /
func main() {
	db := dbconnection.NewDatabase()

	connection, err := db.Open()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	defer db.Close()
	router := gin.Default()

	router.POST("/register", func(ctx *gin.Context) {
		registerQueue(ctx, connection)
	})

	router.GET("/ticket", func(ctx *gin.Context) {
		getTicket(ctx, connection)
	})

	// Start the server
	router.Run(":4444") // Listen on port 4444
}
