package game

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	inp   *bufio.Scanner
}

func (c *CLI) PlayGame() {
	userInput := c.readLine()
	c.store.RecordWin(extractPlayer(userInput))
}

func (c *CLI) readLine() string {
	c.inp.Scan()
	return c.inp.Text()
}

func extractPlayer(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func NewCLi(store PlayerStore, inp io.Reader) *CLI {
	return &CLI{
		store: store,
		inp:   bufio.NewScanner(inp),
	}
}
