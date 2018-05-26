package thread

import (
	"core_backend/user"
	"core_backend/category"

	"gopkg.in/mgo.v2/bson"

	"time"		
)

type Thread struct {
	Id 			bson.ObjectId			`json:"_id" bson:"_id"`
	Title		string					`json:"title" validate:"required"`
	Content		string					`json:"content" validate:"required"`
	Date 		time.Time 				`json:"date"`
	Category 	category.Category		`json:"category" validate:"required"`
	User 		user.UserProtected		`json:"user" validate:"required"`
	Active		bool					`json:"active"`
}

type ThreadUpdate struct {
	Id 			bson.ObjectId			`json:"_id" bson:"_id,omitempty"`
	Title		string					`json:"title" bson:"title,omitempty"`
	Content		string					`json:"content" bson:"content,omitempty"`
	Date 		time.Time 				`json:"date" bson:"date,omitempty"`
	Category 	category.CategoryUpdate	`json:"category" bson:"category,omitempty"`
	User 		user.UserProtected		`json:"user" bson:"user,omitempty"`
	Active		*bool					`json:"active" bson:"active,omitempty"`
}

type Threads []Thread

func NewThread() (thread Thread) {
	thread = Thread{
		Id: bson.NewObjectId(),
		Date: time.Now(),
		Active: true,
	}

	return
}