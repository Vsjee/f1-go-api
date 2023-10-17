package main

import (
	"f1-go-api/configs"
	"f1-go-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//fiber instance
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	//port
	app.Listen("0.0.0.0:" + "6000")
}
