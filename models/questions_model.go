package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Questions struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Question string             `json:"question,omitempty" validate:"required"`
	Answers  []string           `json:"answers,omitempty" validate:"required"`
}
