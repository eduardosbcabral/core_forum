package gender

import (
	"net/http"
	"log"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type GenderController struct {
	GenderRepository GenderRepository
}

func (gc *GenderController) Index(w http.ResponseWriter, r *http.Request) {
	genders, err := gc.GenderRepository.GetGenders()

	if err != nil {
		log.Print("[ERROR] cant find genders")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, genders)
}

func (gc *GenderController) Create(w http.ResponseWriter, r *http.Request) {
	var g Gender

	err := config.DecodeJson(r.Body, &g)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gender, err := gc.GenderRepository.InsertGender(g)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusCreated, gender)
}

func (gc *GenderController) Show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	gender, err := gc.GenderRepository.GetGender(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, gender)
}

func (gc *GenderController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var g Gender

	err := config.DecodeJson(r.Body, &g)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	gender, err := gc.GenderRepository.UpdateGender(id, g)

	if err != nil {
		log.Print("[ERROR] cant find or update gender: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, gender)
}

func (gc *GenderController) Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var g Gender

	err := config.DecodeJson(r.Body, &g)	

	_, err = gc.GenderRepository.DeleteGender(id, g)

	if err != nil {
		log.Print("[ERROR] cant delete gender")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("GENERO DESTRUIDO"))
}