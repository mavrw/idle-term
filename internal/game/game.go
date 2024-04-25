package game

import (
	"fmt"

	"github.com/mavrw/terminally-idle/internal/tui/models"

	tea "github.com/charmbracelet/bubbletea"
)

type GameConfig struct {
	Title     string
	DebugMode bool
}

type gameInstance struct {
	application *tea.Program
	model       models.GameAppModel
}

func NewGame(config GameConfig) *gameInstance {
	instance := &gameInstance{
		model: models.NewGameModel(config.Title),
	}
	instance.model.ToggleDebugMode(config.DebugMode)
	return instance
}

func (g *gameInstance) StartGame() {
	g.start()
}

func (g *gameInstance) start() {
	//? Maybe use this abstraction later to write the bubble tea view to?
	// terminal := terminal.NewDefaultTerminal()

	g.application = tea.NewProgram(g.model)
	if _, err := g.application.Run(); err != nil {
		fmt.Printf("Unable to start the application: %v", err)
	}
}
