package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	godotenv.Load()

	return os.Getenv("MONGOURI")
}
