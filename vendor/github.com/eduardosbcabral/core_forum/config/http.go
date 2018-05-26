package config

import(
	"net/http"
	"encoding/json"
	"log"
	"io"

	"gopkg.in/go-playground/validator.v9"
)

type Response struct {
	StatusCode	int 	`json:"code"`
	Message 	string 	`json:"message"`
}

type DesactivateStruct struct {
	Active bool	`json:"active"`
}

var Responses = map[string]string {
	"success": `Success.`,
	"login": `Logged in.`,
	"bad-json": `Wrong JSON.`,
	"bad-validate": `Wrong data.`,
	"bad-insert": `Can't insert.`,
	"bad-find": `Can't get.`,
	"bad-update": `Can't update.`,
	"bad-destroy": `Can't destroy.`,
	"bad-login": `Wrong username or password.`,
	"created": `Successfully created.`,
	"signup": "Account successfully created.",
	"destroyed": `Successfully destroyed.`,
	"not-found": `Can't found.`,
	"unauthorized": `You're not authorized.`,
}

func BodyValidate(r *http.Request, entity interface{}) interface{} {

	err := DecodeJson(r.Body, entity)

	if err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return Responses["bad-json"]
	}

	err = Validate.Struct(entity)

	if err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)

		var fields []string

		for _, value := range err.(validator.ValidationErrors) {
			fields = append(fields, value.Tag() + ":" + value.Field())
		}
		return fields[len(fields)-1]
	}

	return nil
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

func HttpMessageResponse(w http.ResponseWriter, code int, payload interface{}) {

	r := Response{
		StatusCode: code,
		Message: payload.(string),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}