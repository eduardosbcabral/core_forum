package user

import (
	"net/http"

	"github.com/gorilla/mux"

	"core_backend/config"
)

type UserController struct {
	UserRepository UserRepository
}

func (uu *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := uu.UserRepository.GetUsers()

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return users.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, users)
}

func (uu *UserController) IndexAll(w http.ResponseWriter, r *http.Request) {
	users, err := uu.UserRepository.GetAllUsers()

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return users.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, users)
}

func (uu *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var u User

	err := config.DecodeJson(r.Body, &u)
	
	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}

	user, err := uu.UserRepository.InsertUser(u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't insert user.")
		return
	}

	config.RespondWithJson(w, http.StatusCreated, user)
}

func (uu *UserController) Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := uu.UserRepository.GetUser(username)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't find user.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, user)
}

func (uu *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var u UserLogin

	err := config.DecodeJson(r.Body, &u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}
	
	_, err = uu.UserRepository.Login(u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong username or password.")
		return
	}

	config.RespondWithMessage(w, http.StatusOK, "Logged In.")	
}

func (uu *UserController) Update(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	var u User

	err := config.DecodeJson(r.Body, &u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}
	
	user, err := uu.UserRepository.UpdateUser(username, u)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't update user.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, user)
}

func (uu *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	var ds config.DesactivateStruct

	err := config.DecodeJson(r.Body, &ds)
	
	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	_, err = uu.UserRepository.DeleteUser(username, ds)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't delete user.")
		return
	}

	config.RespondWithMessage(w, http.StatusOK, "User successfully deleted.")
}