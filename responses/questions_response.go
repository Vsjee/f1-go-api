package responses

import "github.com/gofiber/fiber/v2"

type QuestionResponse struct {
	Status  int
	Message string
	Data    *fiber.Map
}
