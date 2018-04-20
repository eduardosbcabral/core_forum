package gender

import (
	"gopkg.in/mgo.v2/bson"
)

type Gender struct {
	Id				bson.ObjectId	`json:"_id" bson:"_id"`
	Description 	string			`json:"description" validate:"required"`
	Active			bool			`json:"active"`
}

type Genders []Gender