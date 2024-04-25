package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/tui/commands"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
	"github.com/mavrw/terminally-idle/internal/tui/messages"
)

type GameAppModel struct {
	title            string
	currentGameState constants.GameState
	states           map[constants.GameState]tea.Model
	debugMode        bool
}

func NewGameModel(title string) GameAppModel {
	return GameAppModel{
		title:            title,
		currentGameState: 0,
		states:           initializeGameStates(),
		debugMode:        false,
	}
}

func (m *GameAppModel) ToggleDebugMode(b bool) {
	m.debugMode = b
}

func (m GameAppModel) Init() tea.Cmd {
	// Set initial state main menu
	return commands.ChangeApplicationState(constants.MAIN_MENU)
}

func (m GameAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Always process quit key message first
	if mes, ok := msg.(tea.KeyMsg); ok && key.Matches(mes, constants.KeyMap.Quit) {
		return m, tea.Quit
	}

	// Keybinding to reset to initial state for debug purposes
	if mes, ok := msg.(tea.KeyMsg); ok && key.Matches(mes, constants.KeyMap.DEBUG_RESET) && m.debugMode {
		return m, commands.ChangeApplicationState(constants.INITIAL_STATE)
	}

	// Pre-process msg
	switch msg := msg.(type) {
	case messages.ChangeStateMsg:
		m.changeGameState(msg.State)
		if m.currentGameState != constants.INITIAL_STATE {
			return m, m.states[m.currentGameState].Init()
		}
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

func (m GameAppModel) View() string {
	if m.currentGameState != constants.INITIAL_STATE {
		if state, ok := m.states[m.currentGameState]; ok {
			return state.View()
		}
	}

	var s string

	if m.debugMode {
		// Header
		s += fmt.Sprintf("%s Debug Menu\n\n\n", m.title)

		s += fmt.Sprintf("GameState: %d\n\n\n\n", m.currentGameState)
	}

	s += "Press CTRL+C to exit..."

	return s
}

func (m *GameAppModel) changeGameState(state constants.GameState) {
	m.currentGameState = state
}

func initializeGameStates() map[constants.GameState]tea.Model {
	return map[constants.GameState]tea.Model{
		constants.MAIN_MENU: NewMainMenu("Main Menu", []MenuOption{
			{
				Label: "Play",
				Cmd:   commands.ChangeApplicationState(constants.IDLE),
			},
			{
				Label: "Settings",
				Cmd:   nil,
			},
			{
				Label: "Exit",
				Cmd:   tea.Quit,
			},
		}),
		constants.IDLE: NewIdleGameModel(),
	}
}
