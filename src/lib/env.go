package lib

import (
	"os"

	"github.com/HidemaruOwO/nuts/log"
	"github.com/joho/godotenv"
)

var (
	PORT             string
	DOMAIN           string
	APP_ENV          string
	VIRTUALSLIME_DIR string
	ISDEBUG          bool
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Warn("Failed loading .env file")
		log.Error(err)
	}

	PORT = os.Getenv("PORT")
	DOMAIN = os.Getenv("DOMAIN")
	APP_ENV = os.Getenv("APP_ENV")
	VIRTUALSLIME_DIR = os.Getenv("VIRTUALSLIME_DIR")
	ISDEBUG = os.Getenv("DEBUG") == "true"
}
