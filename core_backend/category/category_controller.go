package category

import (
	"net/http"
	"log"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	CategoryRepository CategoryRepository
}

func (cc *CategoryController) Index(w http.ResponseWriter, r *http.Request) {
	categories, err := cc.CategoryRepository.GetCategories()

	if err != nil {
		log.Print("[ERROR] cant find categories")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, categories)
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var c Category

	err := config.DecodeJson(r.Body, &c)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	category, err := cc.CategoryRepository.InsertCategory(c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusCreated, category)
}

func (cc *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	category, err := cc.CategoryRepository.GetCategory(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, category)
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c Category

	err := config.DecodeJson(r.Body, &c)

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	category, err := cc.CategoryRepository.UpdateCategory(id, c)

	if err != nil {
		log.Print("[ERROR] cant find or update category: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.RespondWithJson(w, http.StatusOK, category)
}

func (cc *CategoryController) Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c Category

	err := config.DecodeJson(r.Body, &c)	

	if err != nil {
		log.Print("[ERROR] Wrong JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = cc.CategoryRepository.DeleteCategory(id, c)

	if err != nil {
		log.Print("[ERROR] cant delete category")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("CATEGORIA DESTRUIDA"))
}