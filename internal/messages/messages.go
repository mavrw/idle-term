package messages

import "github.com/mavrw/terminally-idle/internal/constants"

type ChangeStateMsg struct {
	State constants.GameState
}
