package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// PORT server port
var (
	PORT      = 0
	SECRETKEY []byte
	DBDRIVER  = ""
	DBURL     = ""
)

// Load server PORT
func Load() {
	var err error
	err = godotenv.Load(".env.production")
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)

	SECRETKEY = []byte(os.Getenv(("API_SCECRET")))
}
