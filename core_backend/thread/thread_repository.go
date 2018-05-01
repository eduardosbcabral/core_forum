package thread

import (
	"core_backend/config"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const docname = "threads"

type ThreadRepository struct{}

func (tr ThreadRepository) GetThreads(categoryId string, result *Threads) (err error) {

	c := config.OpenSession(docname)

	err = c.Find(
			bson.M{
				"active": true, 
				"category._id": bson.ObjectIdHex(categoryId),
			},
		).All(result) 

	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}

	return
}

func (tr ThreadRepository) GetThread(categoryId string, threadId string, result *Thread) (err error) {
	c := config.OpenSession(docname)

	err = c.Find(
				bson.M{
					"_id": bson.ObjectIdHex(threadId),
					"category._id": bson.ObjectIdHex(categoryId),
				},
			).One(result)

	if err != nil {
		return
	}

	return
}