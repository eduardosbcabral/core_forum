package user

import (
	"strings"
	"log"
	
	"core_backend/config"

	"gopkg.in/mgo.v2/bson"
)

const docname = "users"

type UserRepository struct{}

func (ur UserRepository) InsertUser(user *User) (err error) {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	hash, err := config.HashPassword(user.Password)
	user.Password = hash

	if err != nil {
		log.Print("[ERROR] Can't hash password: ", err)
		return
	}

	err = config.Insert(&user, docname)

	if err != nil {		
		log.Print("[ERROR] Can't insert user: ", err)
		return
	}

	return
}

func (ur UserRepository) Login(user UserLogin) (User, bool) {
	c := config.OpenSession(docname)
	u := User{}

	user.Username = strings.ToLower(user.Username)

	err := c.Find(bson.M{"username": user.Username}).One(&u)
	
	if err != nil {
		log.Print("[ERROR] failed to login: ", err)
		return u, false
	}

	match := config.CheckPasswordHash(user.Password, u.Password)

	if !match {
		return u, false
	}

	u.Password = ""

	return u, true

}