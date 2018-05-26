package category

import (
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id				bson.ObjectId	`json:"_id" bson:"_id"`
	Category 		string			`json:"category" validate:"required"`
	Description 	string			`json:"description" validate:"required"`
	Active			bool			`json:"active"`
}


type CategoryUpdate struct {
	Id				bson.ObjectId	`json:"_id" bson:"_id,omitempty"`
	Category 		string			`json:"category" bson:"category,omitempty"`
	Description 	string			`json:"description" bson:"description,omitempty"`
	Active			*bool			`json:"active" bson:"active,omitempty"`
}

type Categories []Category

func NewCategory() (category Category) {
	category = Category{
		Id: bson.NewObjectId(),
		Active: true,
	}

	return
}