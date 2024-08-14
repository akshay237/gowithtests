package main

import (
	"encoding/json"
	"io"
	"log"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(reader io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		log.Println("problem parsing league", err)
	}
	return league, err
}
