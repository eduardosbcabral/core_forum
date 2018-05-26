package gender

import (
	"gopkg.in/mgo.v2/bson"
)

type Gender struct {
	Id				bson.ObjectId	`json:"_id" bson:"_id"`
	Description 	string			`json:"description" validate:"required"`
	Active			bool			`json:"active"`
}

type GenderUpdate struct {
	Id				bson.ObjectId	`json:"_id" bson:"_id,omitempty"`
	Description 	string			`json:"description" bson:"description,omitempty"`
	Active			*bool			`json:"active" bson:"active,omitempty"`
}

type Genders []Gender

func NewGender() (gender Gender) {
	gender = Gender{
		Id: bson.NewObjectId(),
		Active: true,
	}

	return
}