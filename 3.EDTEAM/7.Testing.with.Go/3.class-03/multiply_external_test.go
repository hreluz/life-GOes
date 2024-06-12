package class03_test

import (
	"testing"

	"github.com/hreluz/go-testing-class-3"
)

func TestMultiply(t *testing.T) {
	want := 6
	got := class03.Multiply(2, 3)

	if got != want {
		t.Errorf("It was expected %d, it got %d", want, got)
	}

}
