package config

import(
	"net/http"
	"encoding/json"
	"log"
	"io"
)

type Response struct {
	StatusCode	int 	`json:"code"`
	Message 	string 	`json:"message"`
}

type DesactivateStruct struct {
	Active bool	`json:"active"`
}

func DecodeJson(body io.Reader, entity interface{}) (err error) {
	d := json.NewDecoder(body)
	err = d.Decode(entity)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
		return
	}

	return
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithMessage(w http.ResponseWriter, code int, message string) {
	var e Response
	e.Message = message
	e.StatusCode = code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}