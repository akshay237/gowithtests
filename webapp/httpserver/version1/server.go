package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

// func (s *PlayerServer) GetPlayerScore(name string) int {
// 	if name == "ram" {
// 		return 20
// 	}
// 	if name == "flyod" {
// 		return 10
// 	}
// 	return -1
// }

/*
version - v0
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	if name == "ram" {
		return "20"
	}
	if name == "flyod" {
		return "10"
	}
	return ""
}
*/
