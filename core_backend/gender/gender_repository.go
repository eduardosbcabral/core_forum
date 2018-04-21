package gender

import (
	"core_backend/config"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const docname = "genders"

type GenderRepository struct{}

func (gr GenderRepository) GetGenders() (results Genders, err error) {

	c := config.OpenSession(docname)

	results = Genders{}

	if err = c.Find(bson.M{"active": true}).All(&results); err != nil {
		log.Print("[ERROR] failed to get genders: ", err)
		return
	}

	return
}

func (gr GenderRepository) GetAllGenders() (results Genders, err error) {

	c := config.OpenSession(docname)

	results = Genders{}

	if err = c.Find(bson.M{}).All(&results); err != nil {
		log.Print("[ERROR] failed to get genders: ", err)
		return
	}

	return
}

func (gr GenderRepository) InsertGender(gender Gender) (Gender, error) {

	c := config.OpenSession(docname)

	gender.Id = bson.NewObjectId()
	gender.Active = true

	if err := c.Insert(gender); err != nil {
		log.Print("[ERROR] failed to insert gender: ", err)
		return gender, err
 	}

	return gender, nil

}

func (gr GenderRepository) GetGender(id string) (result Gender, err error) {
	c := config.OpenSession(docname)
	
	result = Gender{}

	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		log.Print("[ERROR] failed to get gender: ", err)
		return
	}

	return
}

func (gr GenderRepository) UpdateGender(id string, gender Gender) (Gender, error) {
	c := config.OpenSession(docname)

	gender.Id = bson.ObjectIdHex(id)

	if err := c.UpdateId(bson.ObjectIdHex(id), gender); err != nil {
		log.Print("[ERROR] failed to update gender: ", err)
		return gender, err
	} 

	return gender, nil
}

func (gr GenderRepository) DeleteGender(id string, entity config.DesactivateStruct) (bool, error) {
	c := config.OpenSession(docname)

	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"active": entity.Active}}); err != nil {
		log.Print("[ERROR] failed to activate or desactivate gender: ", err)
		return false, err
	}

	return true, nil
}