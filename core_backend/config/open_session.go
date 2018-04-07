package config

import (
	"log"
	"gopkg.in/mgo.v2"
)

const server = "localhost:27017"

const dbName = "core_db"

func OpenSession(docName string) *mgo.Collection {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Print("Failed to establish connection to MongoDB Server: ", err)
	}

	return session.DB(dbName).C(docName)
}