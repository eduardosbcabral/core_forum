package user

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserRepository UserRepository
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users := c.UserRepository.GetUsers()
	respondWithJson(w, http.StatusOK, users)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var u User

	err := json.NewDecoder(body).Decode(&u)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.UserRepository.InsertUser(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusCreated, user)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := c.UserRepository.GetUser(username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {

	body := r.Body

	var u User

	err := json.NewDecoder(body).Decode(&u)

	_, err = c.UserRepository.Login(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("LOGADO"))
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	username := mux.Vars(r)["username"]

	var u User

	err := json.NewDecoder(body).Decode(&u)

	user, err := c.UserRepository.UpdateUser(username, u)

	if err != nil {
		log.Print("[ERROR] cant find or update user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

func (c *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	_, err := c.UserRepository.DeleteUser(username)

	if err != nil {
		log.Print("[ERROR] cant delete user")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("USUARIO DESTRUIDO"))
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}