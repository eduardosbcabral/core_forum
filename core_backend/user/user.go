package user

import (
	"core_backend/gender"
)

type User struct {
	Username 	string			`json:"username" validate:"required"` // Username will be the 'ID'. Can't exists two users with the same username
	Email		string			`json:"emaily" validate:"required, email"`
	Image		string			`json:"image" validate:"url"`
	Password	string			`json:"password" validate:"required"`
	Gender 		gender.Gender	`json:"gender" validate:"required"`
}

type Users []User