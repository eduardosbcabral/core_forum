package user

import (
	"strings"

	"core_backend/gender"
	"core_backend/config"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username 	string			`json:"username" validate:"required,used-username"`
	Email		string			`json:"email" validate:"required,email,used-email"`
	Image		string			`json:"image" validate:"omitempty,url"`
	Password	string			`json:"password" validate:"required,password-length"`
	Gender 		gender.Gender	`json:"gender" validate:"required"`
	Admin		bool			`json:"admin"`
	Active		bool			`json:"active"`
}

type UserUpdate struct {
	Username 	string				`json:"username" bson:"username,omitempty" validate:"used-username"`
	Email		string				`json:"email" bson:"email,omitempty" validate:"used-email"`
	Image		string				`json:"image" bson:"image,omitempty" validate:"omitempty,url"`
	Password	string				`json:"password,omitempty" bson:"password,omitempty" validate:"omitempty,password-length"`
	Gender 		gender.GenderUpdate	`json:"gender" bson:"gender,omitempty"`
	Admin		*bool				`json:"admin" bson:"admin,omitempty"`
	Active		*bool				`json:"active" bson:"active,omitempty"`
}

type UserProtected struct {
	Username 	string				`json:"username" validate="required"`
	Email		string				`json:"email" validate="required"`
	Image		string				`json:"image" validate="omitempty,required"`
	Gender 		gender.GenderUpdate	`json:"gender" validate="required"`
	Admin		bool				`json:"admin"`
	Active		bool				`json:"active"`
}

type UserLogin struct {
	Username 	string	`json:"username" validate:"required"`
	Password 	string	`json:"password" validate:"required"`
}

type Users []UserProtected

func NewUser() (user User) {
	user = User{
		Active: true,
	}

	return
}

	
func ValidateUsedUsername(username validator.FieldLevel) bool {
	c := config.OpenSession("users")
	
	u := strings.ToLower(username.Field().String())

	result := User{}

	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}

	return result.Username == ""
}

func ValidateUsedEmail(email validator.FieldLevel) bool {
	c := config.OpenSession("users")
	
	e := strings.ToLower(email.Field().String())

	result := User{}

	if err := c.Find(bson.M{"email": e}).One(&result); err != nil {
	}

	return result.Email == ""
}

func ValidatePasswordLength(password validator.FieldLevel) bool {

	if len(password.Field().String()) < 6 {
		return false
	}

	return true
}