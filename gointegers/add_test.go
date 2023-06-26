package gointegers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add test", func(t *testing.T) {
		got := Add(1, 2)
		want := 3
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
