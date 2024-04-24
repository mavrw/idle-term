package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/constants"
	"github.com/mavrw/terminally-idle/internal/messages"
)

type GameModel struct {
	title            string
	currentGameState constants.GameState
	states           map[constants.GameState]tea.Model
}

func NewGameModel(title string) GameModel {
	return GameModel{
		title:            title,
		currentGameState: 0,
		states: map[constants.GameState]tea.Model{
			constants.MAIN_MENU: NewMainMenu("Main Menu", []MenuOption{"One", "Two", "Three", "Fart"}),
		},
	}
}

func (m GameModel) Init() tea.Cmd {
	// No I/O needed currently, return nil
	return nil
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if mes, ok := msg.(tea.KeyMsg); ok && mes.String() == "ctrl+c" {
		return m, tea.Quit
	}

	// Pre-process msg
	switch msg := msg.(type) {
	case messages.ChangeStateMsg:
		m.ChangeGameState(msg.State)
		fmt.Printf("ChangeStateMsg: %v", msg)
		//? Return m, nil here?
	}

	// Forwards msg to current state model
	if state, ok := m.states[m.currentGameState]; ok {
		s, c := state.Update(msg)
		m.states[m.currentGameState] = s
		return m, c
	}

	// Post-process msg if in initial state
	//! This is not really a good way to do this, but we'll re-visit later
	//! (Famous last words)
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up":
			m.currentGameState++
		case "down":
			if m.currentGameState > constants.INITIAL_STATE {
				m.currentGameState--
			}
		}
	}

	// Return model for rendering with no command
	return m, nil
}

func (m GameModel) View() string {
	if m.currentGameState != constants.INITIAL_STATE {
		if state, ok := m.states[m.currentGameState]; ok {
			return state.View()
		}
	}

	// Header
	s := fmt.Sprintf("%s\n\n\n", m.title)

	s += fmt.Sprintf("GameState: %d\n\n\n\n", m.currentGameState)

	s += "Press CTRL+C to exit..."

	return s
}

func (m *GameModel) ChangeGameState(state constants.GameState) {
	m.currentGameState = state
}
