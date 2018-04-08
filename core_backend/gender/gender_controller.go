package gender

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type GenderController struct {
	GenderRepository GenderRepository
}

func (c *GenderController) Index(w http.ResponseWriter, r *http.Request) {
	genders := c.GenderRepository.GetGenders()

	respondWithJson(w, http.StatusOK, genders)
}

func (c *GenderController) Create(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	var g Gender

	err := json.NewDecoder(body).Decode(&g)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gender, err := c.GenderRepository.InsertGender(g)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusCreated, gender)
}

func (c *GenderController) Show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	gender, err := c.GenderRepository.GetGender(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusOK, gender)
}

func (c *GenderController) Update(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	id := mux.Vars(r)["id"]

	var g Gender

	err := json.NewDecoder(body).Decode(&g)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gender, err := c.GenderRepository.UpdateGender(id, g)

	if err != nil {
		log.Print("[ERROR] cant find or update gender: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respondWithJson(w, http.StatusOK, gender)
}

func (c *GenderController) Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := c.GenderRepository.DeleteGender(id)

	if err != nil {
		log.Print("[ERROR] cant delete gender")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("GENERO DESTRUIDO"))
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}