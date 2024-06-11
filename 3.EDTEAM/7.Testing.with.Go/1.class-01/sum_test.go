package class1

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)

	if got != want {
		t.Logf("Error: It was expected %d, got %d", want, got)
		t.Fail()
	}
}

// go test
