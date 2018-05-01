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
	Posts 		Posts					`json:"posts"`
	Active		bool					`json:"active"`
}

type ThreadUpdate struct {
	Id 			bson.ObjectId			`json:"_id" bson:"_id"`
	Title		string					`json:"title"`
	Content		string					`json:"content"`
	Date 		time.Time 				`json:"date"`
	Category 	category.CategoryUpdate	`json:"category"`
	User 		user.UserProtected		`json:"user"`
	Posts 		Posts					`json:"posts"`
	Active		bool					`json:"active"`
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