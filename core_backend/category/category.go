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
	Id				bson.ObjectId	`json:"_id" bson:"_id"`
	Category 		string			`json:"category"`
	Description 	string			`json:"description"`
	Active			bool			`json:"active"`
}

type Categories []Category

func NewCategory() (category Category) {
	category = Category{
		Id: bson.NewObjectId(),
		Active: true,
	}

	return
}