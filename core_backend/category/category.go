package category

import (
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id			bson.ObjectId	`json:"_id" bson:"_id"`
	Category 	string			`json:"category" validate:"required"`
	Active		bool			`json:"active"`
}

type Categories []Category