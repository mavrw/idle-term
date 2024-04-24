package models

import tea "github.com/charmbracelet/bubbletea"

type IdleGameModel struct {
	currency            float64 // Using floats for now because precision shouldn't be an issue
	incrementRate       float64 // per second
	incrementAmount     float64 // per increment
	incrementMultiplier float64 // to increment
}

func NewIdleGameModel() tea.Model {
	return IdleGameModel{
		currency:            0,
		incrementRate:       1,
		incrementAmount:     0.1,
		incrementMultiplier: 1,
	}
}

func (m IdleGameModel) Init() tea.Cmd {
	return nil
}

func (m IdleGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m IdleGameModel) View() string {
	s := "GAME SCREEN"

	return s
}
