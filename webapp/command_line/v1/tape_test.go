package game

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file: file}
	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newContent, _ := io.ReadAll(file)
	got := string(newContent)
	want := "abc"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
