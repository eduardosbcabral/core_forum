package config

import(
	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func Update(entity interface{}, collection string, ids ...string) (result interface{}, err error) {

	c := OpenSession(collection)

	change := mgo.Change{
        Update: bson.M{"$set": entity},
        ReturnNew: true,
	}

	if collection == "users" {
		ids[0] = strings.ToLower(ids[0])
		_, err = c.Find(bson.M{"username": ids[0]}).Apply(change, &entity)
	} else if collection == "threads" {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(ids[1]), "category._id": bson.ObjectIdHex(ids[0]),}).Apply(change, &entity)
	} else if collection == "posts" {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(ids[2]), "category_id": bson.ObjectIdHex(ids[0]), "thread_id": bson.ObjectIdHex(ids[1]),}).Apply(change, &entity)
	} else {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(ids[0])}).Apply(change, &entity)	
	}

	if err != nil {
		log.Print("[ERROR] failed update entity: ", err)
		return entity, err
	}

	return entity, nil
}

func RemoveFromDB(id string, collection string) (err error) {

	c := OpenSession(collection)

	id = strings.ToLower(id)

	if collection == "users" {
		err = c.Remove(bson.M{"username": id})
	} else {
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	}

	if err != nil {
		log.Print("[ERROR] Can't remove entity: ", err)
		return
	}

	return
}