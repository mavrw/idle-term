package game

import "github.com/mavrw/terminally-idle/internal/terminal"

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) StartGame() {
	g.start()
}

func (g *Game) start() {
	terminal := terminal.NewDefaultTerminal()

	terminal.WriteOutput("Terminally Idle")
}
