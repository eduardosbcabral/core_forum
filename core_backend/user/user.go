package user

import (
	"core_backend/gender"
)

type User struct {
	Username 	string			`json:"username" validate:"required"`
	Email		string			`json:"email" validate:"required,email"`
	Image		string			`json:"image" validate:"omitempty,url"`
	Password	string			`json:"password" validate:"required"`
	Gender 		gender.Gender	`json:"gender" validate:"required"`
	Active		bool			`json:"active"`
}

type UserLogin struct {
	Username 	string	`json:"username" validate:"required"`
	Password 	string	`json:"password" validate:"required"`
}

type Users []User