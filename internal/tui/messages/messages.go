package messages

import "github.com/mavrw/terminally-idle/internal/tui/constants"

type ChangeStateMsg struct {
	State constants.GameState
}
