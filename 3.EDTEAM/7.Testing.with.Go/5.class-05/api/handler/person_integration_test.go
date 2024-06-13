package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hreluz/go-testing-class-5/storage"
	"github.com/labstack/echo/v4"
)

func TestPerson_Create_integration(t *testing.T) {
	data := []byte(`{"name": "Batman", "age":18 }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := newPerson(&store)
	err := p.create(ctx)

	if err != nil {
		t.Errorf("it was not expected error, but it got %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Status code, it was expected %d, it got %d", http.StatusInternalServerError, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)

	if err != nil {
		t.Errorf("It could unmarshal the body: %v", err)
	}

	expectedMessage := "Person was created successfully"
	if resp.Message != expectedMessage {
		t.Errorf("The message was not the expected, it got %q, it was expected %q", expectedMessage, resp.Message)
	}

}
