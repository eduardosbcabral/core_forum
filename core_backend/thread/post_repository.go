package thread

import (
	"core_backend/config"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const docnameP = "posts"

type PostRepository struct{}

func (pr PostRepository) GetPosts(categoryId string, threadId string, result *Posts) (err error) {

	c := config.OpenSession(docnameP)

	err = c.Find(
			bson.M{
				"active": true, 
				"category_id": bson.ObjectIdHex(categoryId),
				"thread_id": bson.ObjectIdHex(threadId),
			},
		).All(result) 

	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}

	return
}

func (pr PostRepository) GetPost(categoryId string, threadId string, postId string, result *Post) (err error) {
	c := config.OpenSession(docnameP)

	err = c.Find(
				bson.M{
					"_id": bson.ObjectIdHex(postId),
					"category_id": bson.ObjectIdHex(categoryId),
					"thread_id": bson.ObjectIdHex(threadId),
				},
			).One(result)

	if err != nil {
		return
	}

	return
}

func (pr PostRepository) CreatePost(categoryId string, threadId string, post *Post) (err error) {
	c := config.OpenSession(docnameP)

	post.Category_Id = bson.ObjectIdHex(categoryId)
	post.Thread_Id = bson.ObjectIdHex(threadId)

	err = c.Insert(post)

	if err != nil {
		return
	}

	return
}