package lib

import (
	"os"

	"github.com/HidemaruOwO/nuts/log"
	"github.com/joho/godotenv"
)

var (
	PORT   string
	DOMAIN string
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Warn("Error loading .env file")
		log.Critical(err)
	}

	PORT = os.Getenv("PORT")
	DOMAIN = os.Getenv("DOMAIN")
}
