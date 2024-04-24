package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/constants"
	"github.com/mavrw/terminally-idle/internal/messages"
)

func ChangeApplicationState(state constants.GameState) tea.Cmd {
	return func() tea.Msg {
		return messages.ChangeStateMsg{
			State: state,
		}
	}
}
