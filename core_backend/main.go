package main

import(
	"fmt"
	"net/http"
	"log"
)

func main() {
	log.Print("[DEBUG] Starting Server...")

	CreateAllRoutes()
	router := Routes

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println("[ERROR] can't stablish server: ", err.Error())
	}
}