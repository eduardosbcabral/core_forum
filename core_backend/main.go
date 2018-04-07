package main

import(
	"core_back_end/user"
	"core_back_end/config"

	"fmt"
	"net/http"
	"log"
)

func main() {
	log.Print("[DEBUG] Starting Server...")

	routes := user.CreateUserRoutes()
	router := config.NewRouter(routes)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println("[ERROR] can't stablish server: ", err.Error())
	}
}