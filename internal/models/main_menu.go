package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/commands"
	"github.com/mavrw/terminally-idle/internal/constants"
)

type MenuOption string

type MainMenuModel struct {
	title   string
	options []MenuOption
}

func NewMainMenu(title string, opts []MenuOption) tea.Model {
	return MainMenuModel{
		title:   title,
		options: opts,
	}
}

func (m MainMenuModel) Init() tea.Cmd {
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Return to the initial state for testing purposes
		case "esc":
			return m, commands.ChangeApplicationState(constants.INITIAL_STATE)
		}
	}

	return m, nil
}

func (m MainMenuModel) View() string {
	s := fmt.Sprintf("%s\n\n\n", m.title)

	for i, option := range m.options {
		s += fmt.Sprintf("\t%d: %s\n", i, option)
	}

	s += "Press esc to exit the main menu..."

	return s
}
