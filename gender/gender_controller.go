package gender

import (
	"net/http"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type GenderController struct {}

const docname = "genders"

func (gc *GenderController) Index(w http.ResponseWriter, r *http.Request) {

	result := Genders{}	

	err := config.FindAllActivated(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (gc *GenderController) IndexAll(w http.ResponseWriter, r *http.Request) {
	
	result := Genders{}	

	err := config.FindAll(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (gc *GenderController) Create(w http.ResponseWriter, r *http.Request) {
	
	gender := NewGender()
	
	err := config.BodyValidate(r, &gender)

	if err != nil {
		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	err = config.Insert(&gender, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-insert"])
		
		return
	}		

	config.HttpResponse(w, http.StatusOK, gender)

	return

}

func (gc *GenderController) Show(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	result := Gender{}

	err := config.FindOne(id, &result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (gc *GenderController) Update(w http.ResponseWriter, r *http.Request) {
	
	id := mux.Vars(r)["id"]
	gender := GenderUpdate{}
	
	err := config.BodyValidate(r, &gender)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	result, err := config.Update(gender, docname, id)

	if err != nil {
		
		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (gc *GenderController) Destroy(w http.ResponseWriter, r *http.Request) {
	
	id := mux.Vars(r)["id"]
	ds := config.DesactivateStruct{}

	err := config.BodyValidate(r, &ds)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	_, err = config.Update(ds, docname, id)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-destroy"])

		return
	}	

	config.HttpMessageResponse(w, http.StatusOK, config.Responses["destroyed"])

	return
}