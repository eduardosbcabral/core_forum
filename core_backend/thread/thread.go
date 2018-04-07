package thread

import (
	"core_back_end/user"
	"time"		
)

type Thread struct {
	Id 			int64				`json:"id"`
	Title		string				`json:"title" validate:"required"`
	Content		string				`json:"content" validate:"required"`
	Date 		time.time 			`json:"date"`
	Category 	category.Category	`json:"category" validate:"required"`
	User 		user.User			`json:"user" validate:"required"`
	Objects 	[]post.Post 		`json:"posts"`
}