package framework

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	e := echo.New()
	ctx := e.NewContext(r, w)

	Get(ctx)

	if w.Code != http.StatusOK {
		t.Errorf("Status code expected: %d, got: %d", http.StatusOK, w.Code)
	}

	t.Log(w.Body.String())

	got := Person{}
	err := json.NewDecoder(w.Body).Decode(&got)

	if err != nil {
		t.Errorf("It could not process json: %v", err)
	}

	want := Person{
		"Batman",
		35,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("It was expected %v, it got %v ", want, got)
	}
}
