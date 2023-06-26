package godepd_injection

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	NewGreet(&buffer, "chris")

	got := buffer.String()
	want := "Hello chris, welcome"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

// Mocking is replace real things you inject with a pretend version that you can control and inspect in your tests

func TestGreetingHandler(t *testing.T) {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetingHandler)))
}
