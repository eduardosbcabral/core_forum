package user

import (
	"strings"
	"log"
	
	"core_backend/config"

	"gopkg.in/mgo.v2/bson"
)

const docname = "users"

type UserRepository struct{}

func (ur UserRepository) Login(user UserLogin) (bool, error) {
	c := config.OpenSession(docname)

	user.Username = strings.ToLower(user.Username)

	if err := c.Find(bson.M{"username": user.Username, "password": user.Password}).One(&User{}); err != nil {
		log.Print("[ERROR] failed to login: ", err)
		return false, err
	}

	return true, nil

}