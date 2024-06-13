package storage

import (
	"testing"

	"github.com/hreluz/go-testing-class-5/model"
)

func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()

	err := m.Create(nil)

	if err == nil {
		t.Error("It was expected an error, it got nil")
	}

	if err != model.ErrPersonCanNotBeNil {
		t.Errorf("It was expected the error %v, it got %v", model.ErrPersonCanNotBeNil, err)
	}
}

func TestCreate_count_elements(t *testing.T) {
	m := NewMemory()
	p := model.Person{Name: "Batman"}
	err := m.Create(&p)

	if err != nil {
		t.Errorf("It was not an error, it got %v", err)
	}

	if m.currentID != 1 {
		t.Errorf("It was expected 1 element, it got: %d ", m.currentID)
	}
}
