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

	threads, err := tc.ThreadRepository.GetThreads(categoryId)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return threads.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, threads)
}

func (tc *ThreadController) IndexAll(w http.ResponseWriter, r *http.Request) {
	threads, err := tc.ThreadRepository.GetAllThreads()

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't return threads.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, threads)
}

func (tc *ThreadController) Create(w http.ResponseWriter, r *http.Request) {

	var t Thread

	err := config.DecodeJson(r.Body, &t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}

	thread, err := tc.ThreadRepository.InsertThread(t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't insert thread.")
		return
	}

	config.RespondWithJson(w, http.StatusCreated, thread)
}

func (tc *ThreadController) Show(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]

	thread, err := tc.ThreadRepository.GetThread(categoryId, threadId)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't find thread.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, thread)
}

func (tc *ThreadController) Update(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]

	var t Thread

	err := config.DecodeJson(r.Body, &t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	err = config.Validate.Struct(t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong data.")
		return	
	}
	
	thread, err := tc.ThreadRepository.UpdateThread(categoryId, threadId, t)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't update thread.")
		return
	}

	config.RespondWithJson(w, http.StatusOK, thread)
}

func (tc *ThreadController) Destroy(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]

	var ds config.DesactivateStruct

	err := config.DecodeJson(r.Body, &ds)	

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Wrong JSON.")
		return
	}

	_, err = tc.ThreadRepository.DeleteThread(categoryId, threadId, ds)

	if err != nil {
		config.RespondWithMessage(w, http.StatusBadRequest, "Can't delete thread.")
		return
	}

	config.RespondWithMessage(w, http.StatusOK, "Thread successfully deleted.")	
}