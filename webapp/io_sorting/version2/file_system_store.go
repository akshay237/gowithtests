package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, io.SeekStart)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}
