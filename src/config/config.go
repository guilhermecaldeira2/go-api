package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Database connection string
	ConnectionString = ""
	// API port
	Port = 0
)

// Initialize database env variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if Port, err = strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}

	ConnectionString = fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER_PASS"),
		os.Getenv("DB_HOST"),
	)
}
