package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct{
    ID  primitive.ObjectID `json:"_id"`
    Status string `json:"status"`
}