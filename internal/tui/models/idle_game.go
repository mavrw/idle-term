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

type ResourceGenerator struct {
	name                 string
	desc                 string
	productionAmount     int     // per production
	productionRate       int     // productions per second
	productionMultiplier float64 // to productionAmount
	quantity             int     // number of generators
}

type IdleGameModel struct {
	currency            float64 // Using floats for now because precision shouldn't be an issue
	incrementRate       int64   // per second
	incrementAmount     float64 // per increment
	incrementMultiplier float64 // to increment
	timer               stopwatch.Model
	startTime           int64
	timeElapsed         int64
	lastIncrementTime   int64
	generators          []ResourceGenerator
	generatorViewPort   viewport.Model
}

func NewIdleGameModel() tea.Model {
	m := IdleGameModel{
		currency:            0,
		incrementRate:       1,
		incrementAmount:     0.1,
		incrementMultiplier: 1,
		timer:               stopwatch.NewWithInterval(time.Microsecond),
		generators: []ResourceGenerator{
			{
				name:                 "Crypto Jacker",
				desc:                 "Machines infected with malware that mines bitcoins using the host's hardware.",
				productionAmount:     1,
				productionRate:       2,
				productionMultiplier: 1,
				quantity:             1,
			},
			{
				name:                 "Ransomware Operations",
				desc:                 "Develop and lease ransomware toolkits to cyber-gangs and receive a cut for passive income.",
				productionAmount:     3,
				productionRate:       1,
				productionMultiplier: 1,
				quantity:             1,
			},
			{
				name:                 "Transaction Leech",
				desc:                 "A worm that burrows into corporate networks and siphons fractions of a cent off each transaction, remaining undetected while leeching in a significant amount of money.",
				productionAmount:     1,
				productionRate:       8,
				productionMultiplier: 1,
				quantity:             1,
			},
			{
				name:                 "Password Hash Cracking",
				desc:                 "Leverage immense computing power for bruteforce database dumps and crack passwords to sell them in batches on the darknet.",
				productionAmount:     12,
				productionRate:       1,
				productionMultiplier: 1,
				quantity:             1,
			},
		},
		generatorViewPort: viewport.New(48, 24),
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

func getStringFromGenerators(rg []ResourceGenerator) string {
	var s string
	for _, gen := range rg {
		s += fmt.Sprintf("%v x%v\t%v/s\n", gen.name, gen.quantity, gen.productionAmount*gen.productionRate)
	}

	return s
}
