package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("couls not listen on port 5000 %v", err)
	}
}
