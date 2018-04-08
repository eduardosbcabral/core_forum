package gender

import (
	"core_backend/config"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const docname = "genders"

type GenderRepository struct{}

func (u GenderRepository) GetGenders() Genders {

	c := config.OpenSession(docname)

	results := Genders{}

	if err := c.Find(bson.M{}).All(&results); err != nil {
		log.Print("Failed to write results: ", err)
	}
	return results
}

func (u GenderRepository) InsertGender(gender Gender) (Gender, error) {

	c := config.OpenSession(docname)

	gender.Id = bson.NewObjectId()

	if err := c.Insert(gender); err != nil {
		return gender, err
	}

	return gender, nil

}

func (u GenderRepository) GetGender(id string) (Gender, error) {
	c := config.OpenSession(docname)
	
	result := Gender{}

	if err := c.FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (u GenderRepository) UpdateGender(id string, gender Gender) (Gender, error) {
	c := config.OpenSession(docname)

	gender.Id = bson.ObjectIdHex(id)

	if err := c.UpdateId(bson.ObjectIdHex(id), gender); err != nil {
		return gender, err
	} 

	return gender, nil
}

func (u GenderRepository) DeleteGender(id string) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Remove(bson.M{"_id": id}); err != nil {
		return false, err
	}

	return true, nil
}