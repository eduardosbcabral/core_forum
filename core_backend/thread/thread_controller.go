package thread

import (
	"net/http"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type ThreadController struct {
	ThreadRepository ThreadRepository
}
	
func (tc *ThreadController) Index(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	result := Threads{}

	err := tc.ThreadRepository.GetThreads(categoryId, &result)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (tc *ThreadController) IndexAll(w http.ResponseWriter, r *http.Request) {
	
	result := Threads{}	

	err := config.FindAll(&result, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (tc *ThreadController) Create(w http.ResponseWriter, r *http.Request) {

	thread := NewThread()

	if !config.BodyValidate(r, &thread) {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])
		
		return
	}

	err := config.Insert(&thread, docname)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-insert"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, thread)

	return
}

func (tc *ThreadController) Show(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	result := Thread{}

	err := tc.ThreadRepository.GetThread(categoryId, threadId, &result)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (tc *ThreadController) Update(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	thread := ThreadUpdate{}

	if !config.BodyValidate(r, &thread) {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])

		return
	}

	result, err := config.Update(thread, docname, categoryId, threadId)

	if err != nil {
		
		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (tc *ThreadController) Destroy(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	ds := config.DesactivateStruct{}

	if !config.BodyValidate(r, &ds) {
		
		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-json"])

		return
	}

	_, err := config.Update(ds, docname, categoryId, threadId)

	if err != nil {

		config.HttpResponse(w, http.StatusBadRequest, config.Responses["bad-destroy"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, config.Responses["destroyed"])

	return
}