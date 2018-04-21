package thread

import (
	"core_backend/user"
	"core_backend/category"

	"gopkg.in/mgo.v2/bson"

	"time"		
)

type Thread struct {
	Id 			bson.ObjectId		`json:"_id" bson:"_id"`
	Title		string				`json:"title" validate:"required"`
	Content		string				`json:"content" validate:"required"`
	Date 		time.Time 			`json:"date"`
	Category 	category.Category	`json:"category" validate:"required"`
	User 		user.User			`json:"user" validate:"required"`
	Objects 	Posts				`json:"posts"`
	Active		bool				`json:"active"`
}

type InactivateThread struct {
	Active		bool				`json:"active"`
}

type Threads []Thread