package user

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"core_backend/config"
)

type UserController struct {
	UserRepository UserRepository
}

func (uu *UserController) Index(w http.ResponseWriter, r *http.Request) {
	
	result := Users{}	

	err := config.FindAllActivated(&result, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) IndexAll(w http.ResponseWriter, r *http.Request) {
	
	result := Users{}	

	err := config.FindAll(&result, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Create(w http.ResponseWriter, r *http.Request) {
	
	user := User{}

	if !config.BodyValidate(r, &user) {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])
		
		return
	}

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	err := config.Insert(&user, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-insert"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, user)

	return
}

func (uu *UserController) Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)

	result := UserProtected{}

	err := config.FindOne(username, &result, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Login(w http.ResponseWriter, r *http.Request) {
	
	ul := UserLogin{}

	if !config.BodyValidate(r, &ul) {
		
		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])

		return
	}

	_, err := uu.UserRepository.Login(ul)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-login"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, config.Responses["login"])

	return
}

func (uu *UserController) Update(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user := UserUpdate{}

	if !config.BodyValidate(r, &user) {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])

		return
	}

	result, err := config.Update(user, docname, username)

	if err != nil {
		
		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	
	username := mux.Vars(r)["username"]
	ds := config.DesactivateStruct{}

	if !config.BodyValidate(r, &ds) {
		
		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])

		return
	}

	_, err := config.Update(ds, docname, username)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-destroy"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, config.Responses["destroyed"])

	return
}