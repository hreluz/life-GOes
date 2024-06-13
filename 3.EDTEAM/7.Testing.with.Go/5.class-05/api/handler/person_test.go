package handler

import (
	"bytes"
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
}
