package env

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Info("Error loading .env file")
	}

	return os.Getenv(key)
}
