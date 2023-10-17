package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Questions struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Question string             `json:"question,omitempty" validate:"required"`
	Answers  []string           `json:"answers,omitempty" validate:"required"`
}
