package main

import(
	"core_backend/config"
)

func main() {
	StartDatabase()
	config.StartValidator()
	StartServer()
}