package main

import (
	"core_backend/user"
	"core_backend/config"
	"github.com/gorilla/mux"
)

var Routes *mux.Router

func CreateAllRoutes() {
	userRoutes := user.CreateUserRoutes()
	/*genderRoutes := gender.CreateGenderRoutes()
	a := append(userRoutes, genderRoutes ...) Exemplo de como fazer todas as rotas da aplicação
	fmt.Printf("AAA: %+v", a)
	*/ 
	
	Routes = config.NewRouter(userRoutes)
}