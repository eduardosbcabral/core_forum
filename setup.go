package main

import(
	"net/http"
	"log"


	"core_backend/config"
	"core_backend/user"
	"core_backend/gender"
	"core_backend/category"
	"core_backend/thread"
	"core_backend/auth"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartDatabase() {
	err := config.ConnectToDatabase()

	if err != nil {
		panic(err)
	}
}

func StartServer() {
	config.ConnectToDatabase()
	config.StartValidator()
	StartCustomValidators()

	c := cors.New(cors.Options{
    	AllowedOrigins: []string{"*"},
    	AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
    	AllowCredentials: true,
    })


 	router := CreateAllRoutes()

 	//Paths that need authorization
 	u := router.PathPrefix("/user").Subrouter()
	u.Methods(http.MethodGet).Path("/{username}").Handler(auth.SetAuthenticatedMiddleware(user.Controller.Show))
	u.Methods(http.MethodPut).Path("/{username}").Handler(auth.SetAuthenticatedMiddleware(user.Controller.Update))
	u.Methods(http.MethodDelete).Path("/{username}").Handler(auth.SetAuthenticatedMiddleware(user.Controller.Destroy))


 	handler := c.Handler(router)

 	log.Print("[DEBUG] Starting Server...")
	if err := http.ListenAndServe(config.SERVER_HOST, handler); err != nil {
		panic(err)
	}
}

func CreateAllRoutes() (routes *mux.Router){
	userRoutes := user.CreateUserRoutes()
	genderRoutes := gender.CreateGenderRoutes()
	categoryRoutes := category.CreateCategoryRoutes()
	threadRoutes := thread.CreateThreadRoutes()
	postRoutes := thread.CreatePostRoutes()
	authRoutes := auth.CreateAuthRoutes()
	appRoutes := append(userRoutes, genderRoutes ...)
	appRoutes = append(appRoutes, categoryRoutes ...)
	appRoutes = append(appRoutes, threadRoutes ...)
	appRoutes = append(appRoutes, postRoutes ...)
	appRoutes = append(appRoutes, authRoutes ...)


	routes = config.NewRouter(appRoutes)

	return routes
}

func StartCustomValidators() {
	config.Validate.RegisterValidation("used-username", user.ValidateUsedUsername)
	config.Validate.RegisterValidation("used-email", user.ValidateUsedEmail)
	config.Validate.RegisterValidation("password-length", user.ValidatePasswordLength)
}