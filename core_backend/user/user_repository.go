package user

import (
	"core_backend/config"
	"log"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

const docname = "users"

type UserRepository struct{}

func (ur UserRepository) GetUsers() (results Users, err error) {

	c := config.OpenSession(docname)

	results = Users{}

	if err = c.Find(bson.M{"active": true}).All(&results); err != nil {
		log.Print("[ERROR] failed to get users: ", err)
		return
	}
	
	return
}

func (ur UserRepository) GetAllUsers() (results Users, err error) {

	c := config.OpenSession(docname)

	results = Users{}

	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Print("[ERROR] failed to get categories: ", err)
		return
	}
	
	return
}

func (ur UserRepository) InsertUser(user User) (User, error) {

	c := config.OpenSession(docname)

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	user.Active = true

	if err := c.Insert(user); err != nil {
		log.Print("[ERROR] failed to insert user: ", err)
		return user, err
	}

	return user, nil

}

func (ur UserRepository) GetUser(username string) (result User, err error) {
	c := config.OpenSession(docname)
	
	username = strings.ToLower(username)

	result = User{}

	if err = c.Find(bson.M{"username": username}).One(&result); err != nil {
		log.Print("[ERROR] failed to get user: ", err)
		return
	}

	return
}

func (ur UserRepository) Login(user UserLogin) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Find(bson.M{"username": user.Username, "password": user.Password}).One(&User{}); err != nil {
		log.Print("[ERROR] failed to login: ", err)
		return false, err
	}

	return true, nil

}

func (ur UserRepository) UpdateUser(username string, user User) (User, error) {
	c := config.OpenSession(docname)

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	if err := c.Update(bson.M{"username": username}, user); err != nil {
		log.Print("[ERROR] failed to update user: ", err)
		return user, err
	} 

	return user, nil
}

func (ur UserRepository) DeleteUser(username string, entity config.DesactivateStruct) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Update(bson.M{"username": username}, bson.M{"$set": bson.M{"active": entity.Active}}); err != nil {
		log.Print("[ERROR] failed to delete user: ", err)
		return false, err
	}


	return true, nil
}