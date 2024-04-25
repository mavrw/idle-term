package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/tui/commands"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
)

type DebugMenuModel struct {
	gameAppModel GameAppModel
	stateToSet   constants.GameState
}

func NewDebugMenuModel(gaModel *GameAppModel) tea.Model {
	return DebugMenuModel{
		gameAppModel: *gaModel,
		stateToSet:   0,
	}
}

func (m DebugMenuModel) Init() tea.Cmd {
	return nil
}

func (m DebugMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	//! This is not really a good way to do this, but we'll re-visit later
	//! (Famous last words)
	//? One day later and I have no clue what that is referring to...
	// TODO: This needs moved to it's own debug_menu model
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch {
		case key.Matches(msg, constants.KeyMap.Up):
			m.stateToSet++
		case key.Matches(msg, constants.KeyMap.Down):
			if m.stateToSet > constants.INITIAL_STATE {
				m.stateToSet--
			}
		case key.Matches(msg, constants.KeyMap.Enter):
			cmds = append(cmds, commands.ChangeApplicationState(m.stateToSet))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m DebugMenuModel) View() string {
	var s string

	// Header
	s += fmt.Sprintf("%s Debug Menu\n\n\n", m.gameAppModel.title)

	s += fmt.Sprintf("GameState: %v\n", m.gameAppModel.currentGameState)
	s += fmt.Sprintf("State to set: %v\n\n\n\n", m.stateToSet)
	s += "Press CTRL+C to exit..."

	return s
}
