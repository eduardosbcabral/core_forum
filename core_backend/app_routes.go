package main

import (
	"core_backend/user"
	"core_backend/gender"
	"core_backend/config"

	"github.com/gorilla/mux"
)

var Routes *mux.Router

func CreateAllRoutes() {
	userRoutes := user.CreateUserRoutes()
	genderRoutes := gender.CreateGenderRoutes()
	appRoutes := append(userRoutes, genderRoutes ...)
	
	Routes = config.NewRouter(appRoutes)
}