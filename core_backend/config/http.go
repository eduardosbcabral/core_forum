package config

import(
	"net/http"
	"encoding/json"
	"log"
	"io"
)

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