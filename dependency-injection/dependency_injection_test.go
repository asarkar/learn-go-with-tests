package dependencyinjection_test

import (
	"bytes"
	dependencyinjection "learn-go-with-tests/dependency-injection"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	err := dependencyinjection.Greet(&buffer, "Chris")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
