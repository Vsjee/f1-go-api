package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Questions struct {
	_id      primitive.ObjectID
	question string
	answers  []string
}
