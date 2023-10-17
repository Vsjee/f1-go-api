package routes

import (
	"f1-go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/questions", controllers.GetQuestions)
}
