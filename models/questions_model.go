package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Questions struct {
	_id      primitive.ObjectID `json:"id,omitempty"`
	question string             `json:"question,omitempty" validate:"required"`
	answers  []string           `json:"answers,omitempty" validate:"required"`
}
