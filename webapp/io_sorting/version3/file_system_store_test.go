package main

import (
	"encoding/json"
	"os"
	"testing"
)

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an err but got one, %v ", err)
	}
}

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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

		store := FileSystemPlayerStore{database: json.NewEncoder(database)}
		store.RecordWin("Ram")

		got := store.GetPlayerScore("Ram")
		want := 1
		assertScore(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})

	t.Run("get sorted league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":20},
			{"Name":"Ram", "Wins":5},
			{"Name":"Shyam", "Wins":15}
		]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		got := store.GetLeague()
		want := League{
			{"Ram", 5},
			{"Cleo", 10},
			{"Shyam", 15},
			{"Chris", 20},
		}
		assertLeague(t, got, want)

		// read again the league
		got1 := store.GetLeague()
		assertLeague(t, got1, want)
	})
}
