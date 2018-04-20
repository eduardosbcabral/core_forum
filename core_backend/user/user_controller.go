package user

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"

	"core_backend/config"
)

type UserController struct {
	UserRepository UserRepository
}

func (uu *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := uu.UserRepository.GetUsers()

	if err != nil {
		log.Print("[ERROR] cant find users")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, users)
}

func (uu *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var u User

	err := config.DecodeJson(r.Body, &u)
	
	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := uu.UserRepository.InsertUser(u)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusCreated, user)
}

func (uu *UserController) Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := uu.UserRepository.GetUser(username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, user)
}

func (uu *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var u User

	err := config.DecodeJson(r.Body, &u)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	_, err = uu.UserRepository.Login(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("LOGADO"))
}

func (uu *UserController) Update(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	var u User

	err := config.DecodeJson(r.Body, &u)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	user, err := uu.UserRepository.UpdateUser(username, u)

	if err != nil {
		log.Print("[ERROR] cant find or update user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, user)
}

func (uu *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	var u User

	err := config.DecodeJson(r.Body, &u)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = uu.UserRepository.DeleteUser(username, u)

	if err != nil {
		log.Print("[ERROR] cant delete user")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("USUARIO DESTRUIDO"))
}