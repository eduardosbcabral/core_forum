package user

import (
	"core_backend/config"
	"log"
	"gopkg.in/mgo.v2/bson"
)

const docname = "users"

type UserRepository struct{}

func (u UserRepository) GetAllUsers() Users {

	c := config.OpenSession(docname)

	results := Users{}

	if err := c.Find(bson.M{}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
	}
	return results
}

func (u UserRepository) InsertUser(user User) (User, error){

	c := config.OpenSession(docname)

	if err := c.Insert(user ,user); err != nil {
		return user, err
	}

	return user, nil

}