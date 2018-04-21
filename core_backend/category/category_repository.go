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
		log.Print("[ERROR] failed to get categories: ", err)
		return
	}

	return
}

func (cr CategoryRepository) GetAllCategories() (result Categories, err error) {

	c := config.OpenSession(docname)

	results := Categories{}

	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Print("[ERROR] failed to get categories: ", err)
		return
	}

	return
}

func (cr CategoryRepository) InsertCategory(category Category) (Category, error) {

	c := config.OpenSession(docname)

	category.Id = bson.NewObjectId()
	category.Active = true

	if err := c.Insert(category); err != nil {
		log.Print("[ERROR] failed to insert category: ", err)
		return category, err
	}

	return category, nil
}

func (cr CategoryRepository) GetCategory(id string) (result Category, err error) {
	c := config.OpenSession(docname)
	
	result = Category{}

	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		log.Print("[ERROR] failed to get category: ", err)
		return
	}

	return
}

func (cr CategoryRepository) UpdateCategory(id string, category Category) (Category, error) {
	c := config.OpenSession(docname)

	category.Id = bson.ObjectIdHex(id)

	if err := c.UpdateId(bson.ObjectIdHex(id), category); err != nil {
		log.Print("[ERROR] failed to update category: ", err)
		return category, err
	} 

	return category, nil
}

func (cr CategoryRepository) DeleteCategory(id string, entity config.DesactivateStruct) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"active": entity.Active}}); err != nil {
		log.Print("[ERROR] failed to activate or desactivate category: ", err)
		return false, err
	}

	return true, nil
}