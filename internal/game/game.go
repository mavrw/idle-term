package game

import (
	"fmt"

	"github.com/mavrw/terminally-idle/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

type gameInstance struct {
	application *tea.Program
	model       models.GameModel
}

func NewGame(title string) *gameInstance {
	return &gameInstance{
		model: models.NewGameModel(title),
	}
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
