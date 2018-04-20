package config

import(
	"log"
	"net/http"
)

//Error abstraction
//NOT WORKING ONLY TESTING
func ErrorHandler(entity interface{}, err error, w http.ResponseWriter, message string) interface{} {

	if err != nil {
		log.Print(message, err)
		w.WriteHeader(http.StatusBadRequest)
		return entity
	}

	return entity
}