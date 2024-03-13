package pkg

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davidspek/paperrockscissors/internal/engine"
)

type Player interface {
	StatsRow() []string
	SetChoice(tea.Msg, engine.Engine) tea.Cmd
	ClearChoice()
	Chosen() engine.Choice
	UpdateStats(result engine.Result)
	stats() stats
}

type HumanPlayer struct {
	Name   string
	Choice engine.Choice
	Stats  stats
}

type stats struct {
	Wins   int
	Losses int
	Ties   int
}

func (s stats) String() []string {
	return []string{fmt.Sprintf("%d", s.Wins), fmt.Sprintf("%d", s.Losses), fmt.Sprintf("%d", s.Ties)}
}

func (p *HumanPlayer) StatsRow() []string {
	s := []string{p.Name}
	s = append(s, p.Stats.String()...)
	return s
}

func (p *HumanPlayer) SetChoice(msg tea.Msg, engine engine.Engine) tea.Cmd {
	choices := engine.Choices()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		for _, c := range choices {
			if msg.String() == c.Key() {
				p.Choice = c
				return nil
			}
		}
	}
	return nil
}

func (p *HumanPlayer) Chosen() engine.Choice {
	return p.Choice
}

func (p *HumanPlayer) ClearChoice() {
	p.Choice = nil
}

func (p *HumanPlayer) UpdateStats(result engine.Result) {
	switch result {
	case engine.Win:
		p.Stats.Wins++
	case engine.Lose:
		p.Stats.Losses++
	case engine.Tie:
		p.Stats.Ties++
	}
}

func (p *HumanPlayer) stats() stats {
	return p.Stats
}

type ComputerPlayer struct {
	Name   string
	Choice engine.Choice
	Stats  stats
}

func (p *ComputerPlayer) StatsRow() []string {
	s := []string{p.Name}
	s = append(s, p.Stats.String()...)
	return s
}

func (p *ComputerPlayer) SetChoice(msg tea.Msg, engine engine.Engine) tea.Cmd {
	choices := engine.Choices()
	p.Choice = choices[rand.Intn(len(choices))]
	return nil
}

func (p *ComputerPlayer) Chosen() engine.Choice {
	return p.Choice
}

func (p *ComputerPlayer) ClearChoice() {
	p.Choice = nil
}

func (p *ComputerPlayer) UpdateStats(result engine.Result) {
	switch result {
	case engine.Win:
		p.Stats.Wins++
	case engine.Lose:
		p.Stats.Losses++
	case engine.Tie:
		p.Stats.Ties++
	}
}

func (p *ComputerPlayer) stats() stats {
	return p.Stats
}
