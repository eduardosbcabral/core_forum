package main

import(
	"github.com/eduardosbcabral/core_forum/config"
)

func main() {
	StartDatabase()
	config.StartValidator()
	StartServer()
}