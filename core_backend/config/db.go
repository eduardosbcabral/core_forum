package config

import (
	"log"
	"gopkg.in/mgo.v2"
)

const server = "localhost:27017"

const dbName = "core_db"

var database *mgo.Database

func ConnectToDatabase() {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Print("Failed to establish connection to MongoDB Server: ", err)
	}

	database = session.DB(dbName)
}

func OpenSession(docName string) *mgo.Collection {
	return database.C(docName)
}