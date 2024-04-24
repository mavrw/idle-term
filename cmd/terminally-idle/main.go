package main

import "github.com/mavrw/terminally-idle/internal/game"

func main() {
	game := game.NewGame()

	game.StartGame()
}
