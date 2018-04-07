package user

import (
	"encoding/json"
	"net/http"
	"log"
)

type UserController struct {
	UserRepository UserRepository
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users := c.UserRepository.GetAllUsers()

	respondWithJson(w, http.StatusOK, users)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var u User

	d := json.NewDecoder(body)
	err := d.Decode(&u)

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

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}