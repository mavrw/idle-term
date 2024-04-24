package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
	title     string
	gameState int
}

func NewGameModel(title string) GameModel {
	return GameModel{
		title:     title,
		gameState: 0,
	}
}

func (m GameModel) Init() tea.Cmd {
	// No I/O needed currently, return nil
	return nil
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.gameState++
		case "down":
			if m.gameState > 0 {
				m.gameState--
			}
		}
	}

	// Return model for rendering with no command
	return m, nil
}

func (m GameModel) View() string {
	// Header
	s := fmt.Sprintf("%s\n\n\n", m.title)

	s += fmt.Sprintf("GameState: %d\n\n\n\n", m.gameState)

	s += "Press CTRL+C to exit..."

	return s
}
