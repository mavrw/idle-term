package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
)

type MenuOption struct {
	Label string
	Cmd   tea.Cmd
}

type MainMenuModel struct {
	title        string
	options      []MenuOption
	optionsIndex int
}

func NewMainMenu(title string, opts []MenuOption) tea.Model {
	return MainMenuModel{
		title:        title,
		options:      opts,
		optionsIndex: 0, //? Is this going to cause a bug later on?
	}
}

func (m MainMenuModel) Init() tea.Cmd {
	// fmt.Print("\007")
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.KeyMap.Up):
			if m.optionsIndex > 0 {
				m.optionsIndex--
			}
		case key.Matches(msg, constants.KeyMap.Down):
			if m.optionsIndex < len(m.options)-1 {
				m.optionsIndex++
			}
		case key.Matches(msg, constants.KeyMap.Enter):
			return m, m.options[m.optionsIndex].Cmd
		}
	}

	return m, nil
}

func (m MainMenuModel) View() string {
	s := fmt.Sprintf("%s\n\n\n", m.title)

	for i, opt := range m.options {
		cursor := " "
		if i == m.optionsIndex {
			cursor = ">"
		}
		s += fmt.Sprintf("%s%s\n", cursor, opt.Label)
	}

	return s
}
