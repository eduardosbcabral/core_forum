package category

import (
	"net/http"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	CategoryRepository CategoryRepository
}

func (cc *CategoryController) Index(w http.ResponseWriter, r *http.Request) {
	categories, err := cc.CategoryRepository.GetCategories()

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return categories.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, categories)
}

func (cc *CategoryController) IndexAll(w http.ResponseWriter, r *http.Request) {
	categories, err := cc.CategoryRepository.GetAllCategories()

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return categories.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, categories)
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var c Category

	err := config.DecodeJson(r.Body, &c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}

	category, err := cc.CategoryRepository.InsertCategory(c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't insert category.")
		return
	}

	config.RespondWithJson(w, http.StatusCreated, category)
}

func (cc *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	category, err := cc.CategoryRepository.GetCategory(id)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't find category.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, category)
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c Category

	err := config.DecodeJson(r.Body, &c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}
	
	category, err := cc.CategoryRepository.UpdateCategory(id, c)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't update category.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, category)
}

func (cc *CategoryController) Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var ds config.DesactivateStruct

	err := config.DecodeJson(r.Body, &ds)	

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	_, err = cc.CategoryRepository.DeleteCategory(id, ds)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't delete category.")
		return
	}

	config.RespondWithMessage(w, http.StatusOK, "Category successfully deleted.")
}