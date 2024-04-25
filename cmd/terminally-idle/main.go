package main

import "github.com/mavrw/terminally-idle/internal/game"

func main() {
	config := game.GameConfig{
		Title:     "Terminally Idle",
		DebugMode: true,
	}
	game := game.NewGame(config)

	game.StartGame()
}
