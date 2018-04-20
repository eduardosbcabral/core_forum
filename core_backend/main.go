package main

import(
	"core_backend/config"
)

func main() {

	err := config.ConnectToDatabase()

	if err != nil {
		panic(err)
	}

	config.StartValidator()
	StartServer()
}