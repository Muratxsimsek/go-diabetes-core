package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Diabetes struct {
	_ID          primitive.ObjectID
	SugarValue   int16
	HungerStatus string
	SugarDate    time.Time
}
