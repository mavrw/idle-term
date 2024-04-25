package models

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mavrw/terminally-idle/internal/tui/commands"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
	"github.com/mavrw/terminally-idle/internal/tui/messages"
)

type gameStateMap map[constants.GameState]tea.Model

type GameAppModel struct {
	title            string
	currentGameState constants.GameState
	states           gameStateMap
	debugMode        bool
}

func NewGameModel(title string) GameAppModel {
	m := GameAppModel{
		title:            title,
		currentGameState: 0,
		debugMode:        false,
	}
	m.initializeGameStates()

	return m
}

func (m *GameAppModel) ToggleDebugMode(b bool) {
	m.debugMode = b
}

func (m GameAppModel) Init() tea.Cmd {
	// Set initial state main menu
	return commands.ChangeApplicationState(constants.MAIN_MENU)
}

func (m GameAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// Always process quit key message first
	if mes, ok := msg.(tea.KeyMsg); ok && key.Matches(mes, constants.KeyMap.Quit) {
		return m, tea.Quit
	}

	// Keybinding to open debug_menu if debug mode enabled
	if mes, ok := msg.(tea.KeyMsg); ok && key.Matches(mes, constants.KeyMap.DEBUG_RESET) && m.debugMode {
		return m, commands.ChangeApplicationState(constants.INITIAL_STATE)
	}

	// Pre-process msg
	switch msg := msg.(type) {
	case messages.ChangeStateMsg:
		m.changeGameState(msg.State)
		//! This might cause issues down the line.
		//! Should mark models with an initialized flag.
		cmd := m.states[m.currentGameState].Init()
		cmds = append(cmds, cmd)
	}

	// Forwards msg to current state model
	if state, ok := m.states[m.currentGameState]; ok {
		s, c := state.Update(msg)
		m.states[m.currentGameState] = s
		cmds = append(cmds, c)
	}

	// Post-process msg

	// Return model for rendering with no command
	return m, tea.Batch(cmds...)
}

func (m GameAppModel) View() string {
	var s string
	if state, ok := m.states[m.currentGameState]; ok {
		s += state.View()
	}

	return lipgloss.NewStyle().MarginLeft(2).Render(s)
}

func (m *GameAppModel) changeGameState(state constants.GameState) {
	m.currentGameState = state
}

func (m *GameAppModel) initializeGameStates() {
	m.states = map[constants.GameState]tea.Model{
		constants.INITIAL_STATE: NewDebugMenuModel(m),
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
