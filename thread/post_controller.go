package thread

import (
	"net/http"

	"core_backend/config"

	"github.com/gorilla/mux"
)

type PostController struct {
	PostRepository PostRepository
}
	
func (pc *PostController) Index(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	result := Posts{}

	err := pc.PostRepository.GetPosts(categoryId, threadId, &result)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (pc *PostController) IndexAll(w http.ResponseWriter, r *http.Request) {
	
	result := Posts{}	

	err := config.FindAll(&result, docnameP)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}	

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (pc *PostController) Create(w http.ResponseWriter, r *http.Request) {

	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]

	post := NewPost()

	err := config.BodyValidate(r, &post)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	err = pc.PostRepository.CreatePost(categoryId, threadId, &post)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-insert"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, post)

	return
}

func (pc *PostController) Show(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	postId := mux.Vars(r)["postId"]
	result := Post{}

	err := pc.PostRepository.GetPost(categoryId, threadId, postId, &result)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["not-found"])

		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (pc *PostController) Update(w http.ResponseWriter, r *http.Request) {
	
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	postId := mux.Vars(r)["postId"]
	post := PostUpdate{}

	err := config.BodyValidate(r, &post)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	result, err := config.Update(post, docnameP, categoryId, threadId, postId)

	if err != nil {
		
		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-update"])
		
		return
	}

	config.HttpResponse(w, http.StatusOK, result)

	return
}

func (pc *PostController) Destroy(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	threadId := mux.Vars(r)["threadId"]
	postId := mux.Vars(r)["postId"]
	ds := config.DesactivateStruct{}

	err := config.BodyValidate(r, &ds)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, err)

		return
	}

	_, err = config.Update(ds, docnameP, categoryId, threadId, postId)

	if err != nil {

		config.HttpMessageResponse(w, http.StatusBadRequest, config.Responses["bad-destroy"])

		return
	}	

	config.HttpMessageResponse(w, http.StatusOK, config.Responses["destroyed"])

	return
}