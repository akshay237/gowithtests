package main

import (
	"fmt"
	game "gowithtests/webapp/command_line/v2"
	"os"

	"github.com/rs/zerolog/log"
)

const dbFileName = "/Users/aks/mine/go/gowithtests/webapp/command_line/v2/game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer close()

	fmt.Println("Let's Player This Weird Game")
	fmt.Println("Type {Name} wins to record a win")
	poker := game.NewCLi(store, os.Stdin)
	poker.PlayGame()
}
