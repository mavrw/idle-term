package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mavrw/terminally-idle/internal/tui/constants"
)

type resourceGenerator string

type IdleGameModel struct {
	currency            float64 // Using floats for now because precision shouldn't be an issue
	incrementRate       int64   // per second
	incrementAmount     float64 // per increment
	incrementMultiplier float64 // to increment
	timer               stopwatch.Model
	startTime           int64
	timeElapsed         int64
	lastIncrementTime   int64
	generators          []resourceGenerator
	generatorViewPort   viewport.Model
}

func NewIdleGameModel() tea.Model {
	m := IdleGameModel{
		currency:            0,
		incrementRate:       1,
		incrementAmount:     0.1,
		incrementMultiplier: 1,
		timer:               stopwatch.NewWithInterval(time.Microsecond),
		generators: []resourceGenerator{
			"Crypto Miners",
			"Ransomware Gangs",
			"Chinese Child Slaves",
		},
		generatorViewPort: viewport.New(16, 24),
	}

	return m
}

func (m IdleGameModel) Init() tea.Cmd {
	cmd := m.timer.Init()
	return cmd
}

func (m IdleGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.KeyMap.Up):
			m.currency++
		case key.Matches(msg, constants.KeyMap.Down):
			m.currency--
		case key.Matches(msg, constants.KeyMap.Enter):
			cmds = append(cmds, m.timer.Toggle())
		}

	case stopwatch.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)

		if m.startTime == 0 {
			m.startTime = time.Now().UnixMilli()
		}
		m.timeElapsed = time.Now().UnixMilli() - m.startTime
		m.update()
		cmds = append(cmds, cmd)

	case stopwatch.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.generatorViewPort.SetContent(getStringFromGenerators(m.generators))

	var cmd tea.Cmd
	m.generatorViewPort, cmd = m.generatorViewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m IdleGameModel) View() string {
	s := "GAME SCREEN\n\n"

	s += fmt.Sprintf("Currency: %v\n", m.currency)
	s += fmt.Sprintf("Timer: %v\n", m.timer)
	s += fmt.Sprintf("Time Elapsed: %v\n\n", m.timeElapsed)
	s += fmt.Sprintf("generators:\n%v\n\n", m.generatorViewPort.View())

	return s
}

func (m *IdleGameModel) update() {
	t := time.Now().UnixMilli() - m.lastIncrementTime
	if t >= 1000/m.incrementRate {
		m.currency += m.incrementAmount
		m.lastIncrementTime = time.Now().UnixMilli()
	}
}

func getStringFromGenerators(rg []resourceGenerator) string {
	var s string
	for _, gen := range rg {
		s += fmt.Sprintf("%v\n", gen)
	}

	return s
}
