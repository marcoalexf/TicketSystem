package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("This script will populate mock data into your database.")
	fmt.Print("Are you sure you want to proceed? (yes/no): ")

	var response string
	fmt.Scanln(&response)
	if response != "yes" {
		fmt.Println("Operation aborted.")
		return
	}

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the connection string from the .env file
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Read the SQL file
	sqlFile := "init_mock_data.sql"
	sqlData, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Error reading SQL file %s: %v", sqlFile, err)
	}

	// Execute the SQL commands
	_, err = db.Exec(string(sqlData))
	if err != nil {
		log.Fatalf("Error executing SQL file: %v", err)
	}

	fmt.Println("Mock data inserted successfully!")
}
