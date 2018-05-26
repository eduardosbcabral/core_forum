package thread

import (
	"time"
	"core_backend/user"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id 			bson.ObjectId		`json:"_id" bson:"_id"`
	Content		string				`json:"content" validate:"required"`
	Date		time.Time 			`json:"date"`
	User		user.UserProtected	`json:"user" validate:"required"`
	Thread_Id	bson.ObjectId		`json:"thread_id"`
	Category_Id	bson.ObjectId 		`json:"category_id"`
	Active		bool				`json:"active"`
}

type PostUpdate struct {
	Id 			bson.ObjectId		`json:"_id" bson:"_id,omitempty"`
	Content		string				`json:"content" bson:"content,omitempty"`
	Date		time.Time 			`json:"date" bson:"date,omitempty"`
	User		user.UserProtected	`json:"user" bson:"user,omitempty"`
	Thread_Id	bson.ObjectId		`json:"thread_id" bson:"thread_id,omitempty"`
	Category_Id	bson.ObjectId		`json:"category_id" bson:"category_id,omitempty"`
	Active 		*bool				`json:"active" bson:"active,omitempty"`
}

type Posts []Post

func NewPost() (post Post) {
	post = Post{
		Id: bson.NewObjectId(),
		Date: time.Now(),
		Active: true,
	}

	return
}