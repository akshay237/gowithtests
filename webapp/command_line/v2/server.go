package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const contentType = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// 1. create new player server intance
	p := new(PlayerServer)

	// 2. assign the store to the player server
	p.store = store

	// 3. get the new router
	router := http.NewServeMux()
	// 3.1 handle the league endpoint
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	// 3.2 handle the players endpoint
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	// 4. assign the router to player server
	p.Handler = router

	// 5. return the player server
	return p
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}
