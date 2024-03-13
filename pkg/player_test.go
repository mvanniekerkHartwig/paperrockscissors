package pkg

import (
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davidspek/paperrockscissors/internal/engine"
)

func TestStatsString(t *testing.T) {
	tests := []struct {
		name        string
		input       stats
		expected    []string
		expectedErr bool
	}{
		{
			"empty stats",
			stats{},
			[]string{"0", "0", "0"},
			false,
		},
		{
			"stats with wins",
			stats{Wins: 1},
			[]string{"1", "0", "0"},
			false,
		},
		{
			"stats with losses",
			stats{Losses: 1},
			[]string{"0", "1", "0"},
			false,
		},
		{
			"stats with ties",
			stats{Ties: 1},
			[]string{"0", "0", "1"},
			false,
		},
		{
			"stats with wins, losses, and ties",
			stats{Wins: 1, Losses: 2, Ties: 3},
			[]string{"1", "2", "3"},
			false,
		},
		{
			"stats in reverse order",
			stats{Wins: 3, Losses: 2, Ties: 1},
			[]string{"1", "2", "3"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.String()
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestPlayerStatsRow(t *testing.T) {
	tests := []struct {
		name        string
		input       Player
		expected    []string
		expectedErr bool
	}{
		{
			"human empty stats",
			&HumanPlayer{Name: "David", Stats: stats{}},
			[]string{"David", "0", "0", "0"},
			false,
		},
		{
			"human all stats",
			&HumanPlayer{Name: "David", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			[]string{"David", "1", "2", "3"},
			false,
		},
		{
			"human stats in reverse order",
			&HumanPlayer{Name: "David", Stats: stats{Wins: 3, Losses: 2, Ties: 1}},
			[]string{"David", "1", "2", "3"},
			true,
		},
		{
			"computer empty stats",
			&ComputerPlayer{Name: "Goliath", Stats: stats{}},
			[]string{"Goliath", "0", "0", "0"},
			false,
		},
		{
			"computer all stats",
			&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			[]string{"Goliath", "1", "2", "3"},
			false,
		},
		{
			"computer stats in reverse order",
			&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 3, Losses: 2, Ties: 1}},
			[]string{"Goliath", "1", "2", "3"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.StatsRow()
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestPlayerChosen(t *testing.T) {
	tests := []struct {
		name        string
		player      Player
		expected    engine.Choice
		expectedErr bool
	}{
		{
			"human rock",
			&HumanPlayer{Choice: &engine.Rock{}},
			&engine.Rock{},
			false,
		},
		{
			"computer paper",
			&ComputerPlayer{Choice: &engine.Paper{}},
			&engine.Paper{},
			false,
		},
		{
			"human no choice",
			&HumanPlayer{},
			nil,
			false,
		},
		{
			"wrong choice",
			&ComputerPlayer{Choice: &engine.Rock{}},
			&engine.Paper{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.player.Chosen()
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestPlayerClearChoice(t *testing.T) {
	tests := []struct {
		name        string
		player      Player
		expected    engine.Choice
		expectedErr bool
	}{
		{
			"human choice selected",
			&HumanPlayer{Choice: &engine.Rock{}},
			nil,
			false,
		},
		{
			"computer no choice selected",
			&HumanPlayer{},
			nil,
			false,
		},
		{
			"computer choice selected",
			&ComputerPlayer{Choice: &engine.Paper{}},
			nil,
			false,
		},
		{
			"computer no choice selected",
			&ComputerPlayer{Choice: &engine.Rock{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player.ClearChoice()
			actual := tt.player.Chosen()
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestPlayerUpdateStats(t *testing.T) {
	tests := []struct {
		name        string
		player      Player
		result      engine.Result
		expected    stats
		expectedErr bool
	}{
		{
			"human win",
			&HumanPlayer{},
			engine.Win,
			stats{Wins: 1},
			false,
		},
		{
			"human loss",
			&HumanPlayer{},
			engine.Lose,
			stats{Losses: 1},
			false,
		},
		{
			"human tie",
			&HumanPlayer{},
			engine.Tie,
			stats{Ties: 1},
			false,
		},
		{
			"computer win",
			&ComputerPlayer{},
			engine.Win,
			stats{Wins: 1},
			false,
		},
		{
			"computer loss",
			&ComputerPlayer{},
			engine.Lose,
			stats{Losses: 1},
			false,
		},
		{
			"computer tie",
			&ComputerPlayer{},
			engine.Tie,
			stats{Ties: 1},
			false,
		},
		{
			"human wrong result",
			&HumanPlayer{Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			"win",
			stats{Wins: 1, Losses: 2, Ties: 3},
			true,
		},
		{
			"computer wrong result",
			&ComputerPlayer{Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			"win",
			stats{Wins: 1, Losses: 2, Ties: 3},
			true,
		},
		{
			"human addition",
			&HumanPlayer{Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			engine.Win,
			stats{Wins: 2, Losses: 2, Ties: 3},
			false,
		},
		{
			"computer addition",
			&ComputerPlayer{Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			engine.Tie,
			stats{Wins: 1, Losses: 2, Ties: 4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player.UpdateStats(tt.result)
			actual := tt.player.stats()
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestHumanPlayerSetChoice(t *testing.T) {
	tests := []struct {
		name   string
		player *HumanPlayer
		engine engine.Engine
	}{
		{
			"rock paper scissors",
			&HumanPlayer{},
			&engine.RockPaperScissors{},
		},
	}
	for _, tt := range tests {
		choices := tt.engine.Choices()
		keyPresses := []string{}
		for _, choice := range choices {
			keyPresses = append(keyPresses, choice.Key())
		}

		for i, choice := range choices {
			expected := choice
			t.Run(tt.name, func(t *testing.T) {
				key := tea.Key{Type: tea.KeyRunes, Runes: []rune(keyPresses[i]), Alt: false}
				tt.player.SetChoice(tea.KeyMsg(key), tt.engine)
				actual := tt.player.Choice
				if !reflect.DeepEqual(actual, expected) {
					t.Errorf("expected %v, got %v", expected, actual)
				}
			})
		}
	}
}

func TestComputerPlayerSetChoice(t *testing.T) {
	tests := []struct {
		name   string
		player *ComputerPlayer
		engine engine.Engine
	}{
		{
			"rock paper scissors",
			&ComputerPlayer{},
			&engine.RockPaperScissors{},
		},
	}
	for _, tt := range tests {
		choices := tt.engine.Choices()
		t.Run(tt.name, func(t *testing.T) {
			tt.player.SetChoice(nil, tt.engine)
			actual := tt.player.Choice
			if !containsChoice(choices, actual) {
				t.Errorf("Choice not a valid option. expected %v, got %v", choices, actual)
			}
		})
	}
}

func containsChoice(choices []engine.Choice, choice engine.Choice) bool {
	for _, c := range choices {
		if c == choice {
			return true
		}
	}
	return false
}
