package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("akki", "")
	want := prefixEnglish + "akki"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestHelloWorld(t *testing.T) {
	t.Run("say hello to akki", func(t *testing.T) {
		got := Hello("akki", "")
		want := prefixEnglish + "akki"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world", func(t *testing.T) {
		got := Hello("", "")
		want := prefixEnglish + "World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in other language", func(t *testing.T) {
		got := Hello("akki", "spanish")
		want := prefixSpanish + "akki"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in french", func(t *testing.T) {
		got := Hello("akki", "french")
		want := "Benjour, akki"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
