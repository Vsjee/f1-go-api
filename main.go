package main

import (
	"f1-go-api/configs"
	"f1-go-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//fiber instance
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//port
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
