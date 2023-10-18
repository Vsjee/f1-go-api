package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return "mongodb+srv://davidfelipehernandez1:rvVAz2hgfAFTSVt1@dbm.h04vkul.mongodb.net/"
}
