package test
/*
import(
	"testing"
)

func TestCreateGenderBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{badDescription": "test"`

	r, _ := TestPostRequest("/gender", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}

func TestCreateGenderWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"badDescription": "test"}`

	r, _ := TestPostRequest("/gender", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}

func TestCreateGenderEmptyDescriptionField(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"description": ""}`

	r, _ := TestPostRequest("/gender", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}
*/