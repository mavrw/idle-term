package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
)

type IdleGameModel struct {
	currency            float64 // Using floats for now because precision shouldn't be an issue
	incrementRate       int64   // per second
	incrementAmount     float64 // per increment
	incrementMultiplier float64 // to increment
	timer               stopwatch.Model
	startTime           int64
	timeElapsed         int64
	lastIncrementTime   int64
}

func NewIdleGameModel() tea.Model {
	m := IdleGameModel{
		currency:            0,
		incrementRate:       1,
		incrementAmount:     1000000,
		incrementMultiplier: 1,
		timer:               stopwatch.NewWithInterval(time.Microsecond),
	}

	return m
}

func (m IdleGameModel) Init() tea.Cmd {
	cmd := m.timer.Init()
	return cmd
}

func (m IdleGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.KeyMap.Up):
			m.currency++
		case key.Matches(msg, constants.KeyMap.Down):
			m.currency--
		case key.Matches(msg, constants.KeyMap.Enter):
			return m, m.timer.Toggle()
		}

	case stopwatch.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)

		if m.startTime == 0 {
			m.startTime = time.Now().UnixMilli()
		}
		m.timeElapsed = time.Now().UnixMilli() - m.startTime
		m.update()

		return m, cmd

	case stopwatch.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m IdleGameModel) View() string {
	s := "GAME SCREEN\n\n"

	s += fmt.Sprintf("Currency: %v\n", m.currency)
	s += fmt.Sprintf("Timer: %v\n", m.timer)
	s += fmt.Sprintf("Time Elapsed: %v\n\n", m.timeElapsed)

	return s
}

func (m *IdleGameModel) update() {
	t := time.Now().UnixMilli() - m.lastIncrementTime
	if t >= 1000/m.incrementRate {
		m.currency += m.incrementAmount
		m.lastIncrementTime = time.Now().UnixMilli()
	}
}
