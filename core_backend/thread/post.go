package thread

import (
	"time"
	"core_backend/user"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id 		bson.ObjectId	`json:"_id" bson:"_id"`
	Content	string			`json:"content" validate:"required"`
	Date	time.Time 		`json:"date"`
	User	user.User		`json:"user" validate:"required"`
	Thread	Thread			`json:"thread" validate:"required"`
}

type Posts []Post