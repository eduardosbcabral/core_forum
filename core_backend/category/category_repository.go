package category

import (
	"core_backend/config"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const docname = "categories"

type CategoryRepository struct{}

func (cr CategoryRepository) GetCategories() (results Categories, err error) {

	c := config.OpenSession(docname)

	results = Categories{}

	if err = c.Find(bson.M{"active": true}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
		return
	}

	return
}

func (cr CategoryRepository) GetAllCategories() (result Categories, err error) {

	c := config.OpenSession(docname)

	results := Categories{}

	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
		return
	}

	return
}

func (cr CategoryRepository) InsertCategory(category Category) (Category, error) {

	c := config.OpenSession(docname)

	category.Id = bson.NewObjectId()
	category.Active = true

	if err := c.Insert(category); err != nil {
		return category, err
	}

	return category, nil
}

func (cr CategoryRepository) GetCategory(id string) (result Category, err error) {
	c := config.OpenSession(docname)
	
	result = Category{}

	if err = c.FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
		return
	}

	return
}

func (cr CategoryRepository) UpdateCategory(id string, category Category) (Category, error) {
	c := config.OpenSession(docname)

	category.Id = bson.ObjectIdHex(id)

	if err := c.UpdateId(bson.ObjectIdHex(id), category); err != nil {
		return category, err
	} 

	return category, nil
}

func (cr CategoryRepository) DeleteCategory(id string, category Category) (bool, error) {
	c := config.OpenSession(docname)

	category.Active = false

	if err := c.UpdateId(bson.ObjectIdHex(id), category); err != nil {
		return false, err
	}

	return true, nil
}