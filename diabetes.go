package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Diabetes struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	SugarValue   int16              `bson:"sugarValue" json:"sugarValue"`
	HungerStatus string             `bson:"hungerStatus" json:"hungerStatus"`
	SugarDate    time.Time          `bson:"sugarDate" json:"sugarDate"`
}
