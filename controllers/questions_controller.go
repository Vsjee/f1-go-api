package controllers

import (
	"context"
	"f1-go-api/configs"
	"f1-go-api/models"
	"f1-go-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var questionsCollection *mongo.Collection = configs.GetCollection(configs.DB, "questions")
var validate = validator.New()

func GetQuestions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var questions []models.Questions
	defer cancel()

	results, err := questionsCollection.Find(ctx, bson.M{})

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuestion models.Questions
		if err = results.Decode(&singleQuestion); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		questions = append(questions, singleQuestion)
	}

	return c.Status(http.StatusOK).JSON(
		responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": questions}},
	)
}
