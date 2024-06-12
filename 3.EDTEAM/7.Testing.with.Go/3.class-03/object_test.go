package class03

import (
	"reflect"
	"testing"
)

func TestDog(t *testing.T) {
	want := &Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Retriever",
		},
	}

	got := &Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Retriever",
		},
	}

	t.Logf("want %p, got %p", want, got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("It was expected %v, it got %v", want, got)
	}
}

//go test -run TestDog
