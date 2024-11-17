package main

import (
	"errors"
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

func generateQrCode(data string) ([]byte, error) {
	// Generate the QR code as a PNG image in memory
	qrCode, err := qrcode.Encode(data, qrcode.Medium, 256) // 256x256 pixels
	if err != nil {
		return nil, errors.New("Cannot generate QRCode for the provided data")
	}
	return qrCode, nil
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not executre query to register new queue"})
		return
	}

	qrCode, err := generateQrCode(distributer.Email + ipAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate QR Code"})
		return
	}

	// Send the QR code image as a response
	ctx.Data(http.StatusOK, "image/png", qrCode)
}

// getQrCode handles a GET request retrieve a qr code that represents a queue
// @Summary      Returns queue
// @Description  Returns queue
// @Tags         qrcode, ticket
// @Produce      json
// @Success      200   {object}  gin.H  "Ticket number for queue"
// @Failure      400   {object}  gin.H  "Bad Request"
// @Router       /ticket [get]
func getQueueQrCode(ctx *gin.Context, connection *pgxpool.Pool) {
	// TODO: Implement the query to fetch and genereate qr code
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
