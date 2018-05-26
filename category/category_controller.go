package category

import (
	"net/http"

	"https://github.com/eduardosbcabral/core_forum/config"

	"github.com/gorilla/mux"
)

type CategoryController struct {}

const docname = "categories"

func (cc *CategoryController) Index(w http.ResponseWriter, r *http.Request) {

	result := Categories{}	

	err := config.FindAllActivated(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (cc *CategoryController) IndexAll(w http.ResponseWriter, r *http.Request) {

	result := Categories{}	

	err := config.FindAll(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	
	category := NewCategory()

	err := config.BodyValidate(r, &category)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	err = config.Insert(&category, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-insert"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, category)

	return
}

func (cc *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	
	id := mux.Vars(r)["id"]
	result := Category{}

	err := config.FindOne(id, &result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	
	id := mux.Vars(r)["id"]
	category := CategoryUpdate{}

	err := config.BodyValidate(r, &category)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	result, err := config.Update(category, docname, id)

	if err != nil {
		
		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (cc *CategoryController) Destroy(w http.ResponseWriter, r *http.Request) {
	
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