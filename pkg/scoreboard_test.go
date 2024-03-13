package pkg

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/table"
)

func TestSortRows(t *testing.T) {
	tests := []struct {
		name        string
		input       []table.Row
		expected    []table.Row
		expectedErr bool
	}{
		{
			"two players with different stats",
			[]table.Row{
				{"David", "1", "2", "3"},
				{"Goliath", "3", "2", "1"},
			},
			[]table.Row{
				{"Goliath", "3", "2", "1"},
				{"David", "1", "2", "3"},
			},
			false,
		},
		{
			"two players with the same stats",
			[]table.Row{
				{"David", "1", "2", "3"},
				{"Goliath", "1", "2", "3"},
			},
			[]table.Row{
				{"David", "1", "2", "3"},
				{"Goliath", "1", "2", "3"},
			},
			false,
		},
		{
			"wrong order",
			[]table.Row{
				{"David", "1", "2", "3"},
				{"Goliath", "3", "2", "1"},
			},
			[]table.Row{
				{"David", "1", "2", "3"},
				{"Goliath", "3", "2", "1"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sortRows(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) && !tt.expectedErr {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestScoreboard_UpdateScores(t *testing.T) {
	type fields struct {
		Model table.Model
	}
	type args struct {
		players []Player
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expected    []table.Row
		expectedErr bool
	}{
		{
			"two players with different stats",
			fields{NewScoreboard().Model},
			args{[]Player{
				&HumanPlayer{Name: "David", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
				&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 3, Losses: 2, Ties: 1}},
			}},
			[]table.Row{
				{"1", "Goliath", "3", "2", "1"},
				{"2", "David", "1", "2", "3"},
			},
			false,
		},
		{
			"two players with the same stats",
			fields{NewScoreboard().Model},
			args{[]Player{
				&HumanPlayer{Name: "David", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
				&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
			}},
			[]table.Row{
				{"1", "David", "1", "2", "3"},
				{"2", "Goliath", "1", "2", "3"},
			},
			false,
		},
		{
			"wrong rank",
			fields{NewScoreboard().Model},
			args{[]Player{
				&HumanPlayer{Name: "David", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
				&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 3, Losses: 2, Ties: 1}},
			}},
			[]table.Row{
				{"2", "Goliath", "3", "2", "1"},
				{"1", "David", "1", "2", "3"},
			},
			true,
		},
		{
			"wrong order",
			fields{NewScoreboard().Model},
			args{[]Player{
				&HumanPlayer{Name: "David", Stats: stats{Wins: 1, Losses: 2, Ties: 3}},
				&ComputerPlayer{Name: "Goliath", Stats: stats{Wins: 3, Losses: 2, Ties: 1}},
			}},
			[]table.Row{
				{"2", "David", "1", "2", "3"},
				{"1", "Goliath", "3", "2", "1"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scoreboard{
				Model: tt.fields.Model,
			}
			s.UpdateScores(nil, tt.args.players)
			if !reflect.DeepEqual(s.Rows(), tt.expected) && !tt.expectedErr {
				t.Errorf("Expected %v, got %v", tt.expected, s.Rows())
			}
		})
	}
}
