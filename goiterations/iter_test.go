package goiterations

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaa"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
