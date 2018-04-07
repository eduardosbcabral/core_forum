package user

import (
	"core_back_end/gender"
)

type User struct {
	Username 	string			`bson:"_id" validate:"required"`
	Email		string			`json:"email" validate:"required,email"`
	Image		string			`json:"image" validate:"url"`
	Password	string			`json:"password" validate:"required"`
	Gender 		gender.Gender	`json:"gender" validate:"required"`
}

type Users []User