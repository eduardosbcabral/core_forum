package config

import (
	"log"
	"reflect"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/mitchellh/mapstructure"
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



func FindAllActivated(entity interface{}, collection string) (err error) {

	c := OpenSession(collection)

	err = c.Find(bson.M{"active": true}).All(entity)

	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}

	return
}

func FindAll(entity interface{}, collection string) (err error) {

	c := OpenSession(collection)

	err = c.Find(nil).All(entity)

	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}

	return
}

func FindOne(id string, result interface{}, collection string) (err error) {
	
	c := OpenSession(collection)
	
	if collection == "users" {
		err = c.Find(bson.M{"username": id}).One(result)
	} else {
		err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(result)
	}

	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}

	return

}

func Insert(entity interface{}, collection string) (err error){
	
	c := OpenSession(collection)

	err = c.Insert(entity)

	if err != nil {
		log.Print("[ERROR] failed to insert entity: ", err)
		return
	}

	return
}


// Generic function to do partial update in database
func Update(entity interface{}, collection string, ids ...string) (interface{}, error) {

	c := OpenSession(collection)

	rids := make([]reflect.Value, len(ids))
    for i, a := range ids {
        rids[i] = reflect.ValueOf(a)
    }

	m := make(map[string]interface{})

	t := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			if t.Field(i).Name != "Id" {
				field := strings.ToLower(t.Field(i).Name)
				m[field] = v.Field(i).Interface()
			}
		}
	}

	update := bson.M{}

	err := mapstructure.Decode(m, &update)

	if err != nil {
		log.Print("[ERROR] Can't map entity: ", err)
		return entity, err
	}

	change := mgo.Change{
        Update: bson.M{"$set": update},
        ReturnNew: true,
	}

	if collection == "users" {
		_, err = c.Find(bson.M{"username": ids[0]}).Apply(change, &entity)
	} else if collection == "threads" {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(rids[1].String()), "category._id": bson.ObjectIdHex(rids[0].String()),}).Apply(change, &entity)
	} else {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(rids[0].String())}).Apply(change, &entity)	
	}

	if err != nil {
		log.Print("[ERROR] failed update entity: ", err)
		return entity, err
	}

	return entity, nil
}

