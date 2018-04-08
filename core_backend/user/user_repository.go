package user

import (
	"core_backend/config"
	"log"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

const docname = "users"

type UserRepository struct{}

func (u UserRepository) GetUsers() Users {

	c := config.OpenSession(docname)

	results := Users{}

	if err := c.Find(bson.M{}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
	}
	return results
}

func (u UserRepository) InsertUser(user User) (User, error) {

	c := config.OpenSession(docname)

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	if err := c.Insert(user); err != nil {
		return user, err
	}

	return user, nil

}

func (u UserRepository) GetUser(username string) (User, error) {
	c := config.OpenSession(docname)
	
	username = strings.ToLower(username)

	result := User{}

	if err := c.Find(bson.M{"username": username}).One(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (u UserRepository) Login(user User) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Find(bson.M{"username": user.Username, "password": user.Password}).One(&User{}); err != nil {
		return false, err
	}

	return true, nil

}

func (u UserRepository) UpdateUser(username string, user User) (User, error) {
	c := config.OpenSession(docname)

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	if err := c.Update(bson.M{"username": username}, user); err != nil {
		return user, err
	} 

	return user, nil
}

func (u UserRepository) DeleteUser(username string) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Remove(bson.M{"username": username}); err != nil {
		return false, err
	}

	return true, nil
}