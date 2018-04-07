package user

import (
	"core_back_end/config"
	"log"
)

const docname = "users"

type UserRepository struct{}

func (u UserRepository) GetAllUsers() Users {

	c := config.OpenSession(docname)

	results := Users{}

	if err := c.Find(nil).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
	}
	return results
}

func (u UserRepository) InsertUser(user User) bool {

	c := config.OpenSession(docname)

	if err := c.Insert(user); err != nil {
		log.Fatal("Failed to add user: ", err)
		return false
	}

	return true

}