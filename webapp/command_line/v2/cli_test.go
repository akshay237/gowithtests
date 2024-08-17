package game_test

import (
	game "gowithtests/webapp/command_line/v2"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record Mahadev win from user input", func(t *testing.T) {
		inp := strings.NewReader("Mahadev wins\n")
		playerStore := &game.StubPlayerStore{}
		cli := game.NewCLi(playerStore, inp)
		cli.PlayGame()
		game.AssertPlayerWin(t, playerStore, "Mahadev")
	})

	t.Run("record Ram win from user input", func(t *testing.T) {
		inp := strings.NewReader("Ram wins\n")
		playerStore := &game.StubPlayerStore{}
		cli := game.NewCLi(playerStore, inp)
		cli.PlayGame()
		game.AssertPlayerWin(t, playerStore, "Ram")
	})
}
