package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hreluz/go-testing-class-5/model"
	"github.com/hreluz/go-testing-class-5/storage"
	"github.com/labstack/echo/v4"
)

func TestPerson_Create_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})

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
		t.Errorf("Status code, it was expected %d, it got %d", http.StatusCreated, w.Code)
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

	cleanData(t)
}

type responsePerson struct {
	MessageType string         `json:"message_type"`
	Message     string         `json:"message"`
	Data        []model.Person `json:"data"`
}

func TestPerson_GetAll_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})

	insertData(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)

	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := newPerson(&store)
	err := p.getAll(ctx)

	if err != nil {
		t.Errorf("it was not expected error, but it got %v", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code, it was expected %d, it got %d", http.StatusOK, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("It could not marshal: %v", err)
	}

	lenPersonsExpected := 2
	lePersonsGot := len(resp.Data)

	if lePersonsGot != lenPersonsExpected {
		t.Errorf("It was expecte %d people, it got %d people", lenPersonsExpected, lePersonsGot)
	}

	firstName1 := "Bats"
	if resp.Data[0].Name != firstName1 {
		t.Errorf("it got %q, it got %q", firstName1, resp.Data[0].Name)
	}

}

func cleanData(t *testing.T) {
	store := storage.NewPsql()
	defer store.CloseDB()

	err := store.DeleteAll()

	if err != nil {
		t.Fatalf("Table could not be truncated: %v", err)
	}
}

func insertData(t *testing.T) {

	store := storage.NewPsql()
	defer store.CloseDB()

	err := store.Create(&model.Person{Name: "Bats", Age: 35})

	if err != nil {
		t.Fatalf("it could register the person %v", err)
	}

	err = store.Create(&model.Person{Name: "Joker", Age: 99})

	if err != nil {
		t.Fatalf("it could register the person %v", err)
	}
}
