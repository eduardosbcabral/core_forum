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
	log.Println(users)

	w.Write([]byte("USUARIO"))
	return
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var u User

	d := json.NewDecoder(body)
	err := d.Decode(&u)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
	}

	success := c.UserRepository.InsertUser(u)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("usuario adicionado"))
	return
}