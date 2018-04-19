package main

import(
	"fmt"
	"net/http"
	"log"

	"core_backend/config"
	"github.com/gorilla/handlers"

)

func main() {
	log.Print("[DEBUG] Starting Server...")

	config.ConnectToDatabase()

	CreateAllRoutes()
	router := Routes

	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	if err := http.ListenAndServe("localhost:8080", handlers.CORS(allowedOrigins, allowedMethods)(router)); err != nil {
		fmt.Println("[ERROR] can't stablish server: ", err.Error())
	}
}