package gender

import (
	"net/http"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type GenderController struct {
	GenderRepository GenderRepository
}

func (gc *GenderController) Index(w http.ResponseWriter, r *http.Request) {
	genders, err := gc.GenderRepository.GetGenders()

	if err != nil {		
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return genders.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, genders)
}

func (gc *GenderController) IndexAll(w http.ResponseWriter, r *http.Request) {
	genders, err := gc.GenderRepository.GetAllGenders()

	if err != nil {		
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return genders.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, genders)
}

func (gc *GenderController) Create(w http.ResponseWriter, r *http.Request) {
	var g Gender

	err := config.DecodeJson(r.Body, &g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}

	gender, err := gc.GenderRepository.InsertGender(g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't insert gender.")
		return
	}

	config.RespondWithJson(w, http.StatusCreated, gender)
}

func (gc *GenderController) Show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	gender, err := gc.GenderRepository.GetGender(id)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't find gender.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, gender)
}

func (gc *GenderController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var g Gender

	err := config.DecodeJson(r.Body, &g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}
	
	gender, err := gc.GenderRepository.UpdateGender(id, g)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't update gender.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, gender)
}

func (gc *GenderController) Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var ds config.DesactivateStruct

	err := config.DecodeJson(r.Body, &ds)	

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	_, err = gc.GenderRepository.DeleteGender(id, ds)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't delete gender.")
		return
	}

	config.RespondWithMessage(w, http.StatusOK, "Gender successfully deleted.")	
}