package post

import (
	"time"
	"core_back_end/user"
	"core_back_end/thread"
)

type Post struct {
	Id 		int64			`json:"id"`
	Content	string			`json:"content" validate:"required"`
	Date	time.Time 		`json:"date"`
	User	user.User		`json:"user" validate:"required"`
	Thread	thread.Thread	`json:"thread" validate:"required"`
}