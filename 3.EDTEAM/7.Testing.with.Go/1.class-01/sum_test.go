package class1

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	t.Logf("want value: %d\n", want)
	got := Add(2, 3)
	t.Logf("got value: %d\n", got)

	if got != want {
		t.Logf("Error: It was expected %d, got %d", want, got)
		t.Fail()
	}
	t.Log("Add test finished")
}

func TestAddMultiple(t *testing.T) {
	want := 6
	got := AddMultiple(2, 3, 1)

	if got != want {
		t.Errorf("Error: It was expected %d, got %d", want, got)
	}
}

// Run tests
// ------------
// go test
// go test -run AddM
// go test -run ple$

// Show Logs
// -------------
// go test -v
// go test -v -run AddM
