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

var Responses = map[string]string {
	"success": `{"code": 200, "message": "Success."}`,
	"login": `{"code": 200, "message": "Logged in."}`,
	"bad-json": `{"code": 400, "message": "Wrong JSON."}`,
	"bad-validate": `{"code": 400, "message": "Wrong data."}`,
	"bad-insert": `{"code": 400, "message": "Can't insert."}`,
	"bad-find": `{"code": 400, "message": "Can't get."}`,
	"bad-update": `{"code": 400, "message": "Can't update."}`,
	"bad-destroy": `{"code": 400, "message": "Can't destroy."}`,
	"bad-login": `{"code": 400, "message": "Username or password wrong."}`,
	"created": `{"code": 201, "message": "Successfully created."}`,
	"destroyed": `{"code": 201, "message": "Successfully destroyed."}`,
	"not-found": `{"code": 400, "message": "Can't found."}`,
	"unauthorized": `{"code": 401, "message": "You're not authorized."}`,
}

func BodyValidate(r *http.Request, entity interface{}) bool {

	err := DecodeJson(r.Body, entity)

	if err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return false
	}

	err = Validate.Struct(entity)

	if err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		return false
	}

	return true
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

func HttpResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}