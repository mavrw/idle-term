package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

type GameState int

const (
	INITIAL_STATE GameState = iota
	MAIN_MENU
	IDLE
	SHOP
)

type keyMap struct {
	Up          key.Binding
	Down        key.Binding
	Left        key.Binding
	Right       key.Binding
	Enter       key.Binding
	Back        key.Binding
	Numbers     key.Binding
	Quit        key.Binding
	DEBUG_RESET key.Binding
}

var KeyMap = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("up", "up arrow"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("down", "down arrow"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "left arrow"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", "right arrow"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding( //? Use for menu tho???
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Numbers: key.NewBinding(
		key.WithKeys("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"),
		key.WithHelp("0-9", "numbers"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	DEBUG_RESET: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "DEBUG RESET"),
	),
}
