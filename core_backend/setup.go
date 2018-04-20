package main

import(
	"net/http"
	"log"

	"core_backend/config"
	"core_backend/user"
	"core_backend/gender"
	"core_backend/category"


	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func StartServer() {
	config.ConnectToDatabase()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

 	router := CreateAllRoutes()

 	log.Print("[DEBUG] Starting Server...")
	if err := http.ListenAndServe(config.SERVER_HOST, handlers.CORS(allowedOrigins, allowedMethods)(router)); err != nil {
		panic(err)
	}
}

func CreateAllRoutes() (routes *mux.Router){
	userRoutes := user.CreateUserRoutes()
	genderRoutes := gender.CreateGenderRoutes()
	categoryRoutes := category.CreateCategoryRoutes()
	appRoutes := append(userRoutes, genderRoutes ...)
	appRoutes = append(appRoutes, categoryRoutes ...)

	routes = config.NewRouter(appRoutes)

	return routes
}