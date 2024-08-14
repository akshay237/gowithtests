package main

import (
	"io"
	"os"
	"testing"
)

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 20},
		}
		assertLeague(t, got, want)

		// this will fail as read do byte by byte and reached EOF
		got1 := store.GetLeague()
		want1 := []Player{
			{"Cleo", 10},
			{"Chris", 20},
		}
		assertLeague(t, got1, want1)
	})

	t.Run("get player's score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}
		got := store.GetPlayerScore("Chris")
		want := 20
		assertScore(t, got, want)
	})

	t.Run("store wins for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 21
		assertScore(t, got, want)
	})

	t.Run("store Wins for new Players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}
		store.RecordWin("Ram")

		got := store.GetPlayerScore("Ram")
		want := 1
		assertScore(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	tempfile, err := os.CreateTemp("./", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempfile.Write([]byte(initialData))
	removeFile := func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}

	return tempfile, removeFile
}
