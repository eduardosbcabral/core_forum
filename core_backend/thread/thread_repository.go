package thread

import (
	"core_backend/config"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const docname = "threads"

type ThreadRepository struct{}

func (tr ThreadRepository) GetThreads(categoryId string) (results Threads, err error) {

	c := config.OpenSession(docname)

	results = Threads{}

	if err = c.Find(
		bson.M{"active": true, 
		"category._id": bson.ObjectIdHex(categoryId)}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
		return
	}

	return
}

func (tr ThreadRepository) GetAllThreads() (result Threads, err error) {

	c := config.OpenSession(docname)

	results := Threads{}

	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
		return
	}

	return
}

func (tr ThreadRepository) InsertThread(thread Thread) (Thread, error) {

	c := config.OpenSession(docname)

	thread.Id = bson.NewObjectId()
	thread.Active = true
	thread.Date = time.Now()

	if err := c.Insert(thread); err != nil {
		return thread, err
	}

	return thread, nil
}

func (tr ThreadRepository) GetThread(categoryId string, threadId string) (result Thread, err error) {
	c := config.OpenSession(docname)

	result = Thread{}

	if err = c.Find(
		bson.M{ "_id": bson.ObjectIdHex(threadId),
		"category._id": bson.ObjectIdHex(categoryId)},
		).One(&result); err != nil {
		return
	}

	return
}

func (tr ThreadRepository) UpdateThread(categoryId string, threadId string, thread Thread) (Thread, error) {
	c := config.OpenSession(docname)

	thread.Id = bson.ObjectIdHex(threadId)

	if err := c.Update(
		bson.M{"_id": bson.ObjectIdHex(threadId), 
		"category._id": bson.ObjectIdHex(categoryId)}, thread); err != nil {
		return thread, err
	} 

	return thread, nil
}

func (tr ThreadRepository) DeleteThread(categoryId string, threadId string, entity config.DesactivateStruct) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Update(
		bson.M{"_id": bson.ObjectIdHex(threadId), 
		"category._id": bson.ObjectIdHex(categoryId)}, bson.M{"$set": bson.M{"active": entity.Active}}); err != nil {
		return false, err
	} 

	return true, nil
}