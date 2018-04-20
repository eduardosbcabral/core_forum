package config

import (
	"log"
	"gopkg.in/mgo.v2"
)

var database *mgo.Database

func ConnectToDatabase() (err error) {
	session, err := mgo.Dial(MONGO_HOST)
	if err != nil {
		log.Print("Failed to establish connection to MongoDB Server: ", err)
		return
	}

	database = session.DB(DATABASE_NAME)

	return
}

func OpenSession(docName string) *mgo.Collection {
	return database.C(docName)
}