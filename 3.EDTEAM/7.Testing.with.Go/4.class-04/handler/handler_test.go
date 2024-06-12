package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	Get(w, r)

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

// go test -run TestGet
// go test -v -run TestGet
