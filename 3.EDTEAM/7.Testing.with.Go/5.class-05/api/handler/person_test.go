package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPerson_Create_wrong_structure(t *testing.T) {

	data := []byte(`{"name": 123, "age":18 }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOk{})
	err := p.create(ctx)

	if err != nil {
		t.Errorf("it was not expected error, but it got %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status code, it was expected %d, it got %d", http.StatusBadRequest, w.Code)
	}

	// t.Logf("Body %v", w.Body.String())
	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)

	if err != nil {
		t.Errorf("It could unmarshal the body: %v", err)
	}

	expectedMessage := "Person does not have the correct structure"
	if resp.Message != expectedMessage {
		t.Errorf("The message was not the expected, it got %q, it was expected %q", expectedMessage, resp.Message)
	}

}
