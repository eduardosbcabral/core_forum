package test
/*
import(
	"testing"
)

// BAD TESTS
func TestCreateCategoryBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badCategory": "test", "badDescription": "test"`

	r, _ := TestPostRequest("/category", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}

func TestCreateCategoryWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"badCategory": "test", "badDescription": "test"}`

	r, _ := TestPostRequest("/category", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}

func TestCreateCategoryEmptyCategoryField(t *testing.T) {

	expected := `{"code":400,"message":"required:Category"}`
	request := `{"category": "", "description": "test"}`

	r, _ := TestPostRequest("/category", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}


func TestCreateCategoryEmptyDescriptionField(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"category": "test", "description": ""}`

	r, _ := TestPostRequest("/category", HEADER_REQUEST_JSON, request)

	resp := TestReaderToString(r.Body)

	if resp != expected {
		t.Errorf("The http response was: %v Expected: %v", resp, expected)
	}

}

// OK TESTS

*/