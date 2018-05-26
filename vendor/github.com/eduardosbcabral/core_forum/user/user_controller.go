package user

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/eduardosbcabral/core_forum/config"
	"github.com/eduardosbcabral/core_forum/auth"
)

type UserController struct {
	UserRepository UserRepository
}

func (uu *UserController) Index(w http.ResponseWriter, r *http.Request) {
	
	result := Users{}	

	err := config.FindAllActivated(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) IndexAll(w http.ResponseWriter, r *http.Request) {
	
	result := Users{}	

	err := config.FindAll(&result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Login(w http.ResponseWriter, r *http.Request) {
	
	ul := UserLogin{}

	err := config.BodyValidate(r, &ul)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	user, status := uu.UserRepository.Login(ul)

	if !status {

		config.HttpMessageResponse(w, http.StatusUnauthorized, config.Responses["bad-login"])

		return
	}	

	token := auth.GenerateJWT(user)


	w.Header().Add("Authorization", "Bearer " + token)
	config.HttpMessageResponse(w, http.StatusOK, config.Responses["login"])

	return
}

func (uu *UserController) Create(w http.ResponseWriter, r *http.Request) {
	
	user := NewUser()

	err := config.BodyValidate(r, &user)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	err = uu.UserRepository.InsertUser(&user)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)
		
		return
	}

	config.HttpMessageResponse(w, http.StatusOK, config.Responses["signup"])

	return
}

func (uu *UserController) Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	username = strings.ToLower(username)

	result := UserProtected{}

	err := config.FindOne(username, &result, docname)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Update(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user := UserUpdate{}

	err := config.BodyValidate(r, &user)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	result, err := config.Update(user, docname, username)

	if err != nil {
		
		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (uu *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	
	username := mux.Vars(r)["username"]
	ds := config.DesactivateStruct{}

	err := config.BodyValidate(r, &ds)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	_, err = config.Update(ds, docname, username)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-destroy"])

		return
	}	

	config.HttpMessageResponse(w, http.StatusOK, config.Responses["destroyed"])

	return
}