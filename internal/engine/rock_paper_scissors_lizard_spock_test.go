package engine

import "testing"

func TestRPSLSChoices(t *testing.T) {
	rpsls := RockPaperScissorsLizardSpock{}
	choices := rpsls.Choices()
	if len(choices) != 5 {
		t.Errorf("expected 5 choices, got %d", len(choices))
	}
	if _, ok := choices[0].(*Rock); !ok {
		t.Errorf("expected Rock, got %T", choices[0])
	}
	if _, ok := choices[1].(*Paper); !ok {
		t.Errorf("expected Paper, got %T", choices[1])
	}
	if _, ok := choices[2].(*Scissors); !ok {
		t.Errorf("expected Scissors, got %T", choices[2])
	}
	if _, ok := choices[3].(*Lizard); !ok {
		t.Errorf("expected Lizard, got %T", choices[3])
	}
	if _, ok := choices[4].(*Spock); !ok {
		t.Errorf("expected Spock, got %T", choices[4])
	}
}

func TestRPSLSResult(t *testing.T) {
	tests := []struct {
		name     string
		rpsls    RockPaperScissorsLizardSpock
		c1, c2   Choice
		expected Result
	}{
		{
			"rock vs rock",
			RockPaperScissorsLizardSpock{},
			&Rock{},
			&Rock{},
			Tie,
		},
		{
			"rock vs paper",
			RockPaperScissorsLizardSpock{},
			&Rock{},
			&Paper{},
			Lose,
		},
		{
			"rock vs scissors",
			RockPaperScissorsLizardSpock{},
			&Rock{},
			&Scissors{},
			Win,
		},
		{
			"rock vs lizard",
			RockPaperScissorsLizardSpock{},
			&Rock{},
			&Lizard{},
			Win,
		},
		{
			"rock vs spock",
			RockPaperScissorsLizardSpock{},
			&Rock{},
			&Spock{},
			Lose,
		},
		{
			"paper vs rock",
			RockPaperScissorsLizardSpock{},
			&Paper{},
			&Rock{},
			Win,
		},
		{
			"paper vs paper",
			RockPaperScissorsLizardSpock{},
			&Paper{},
			&Paper{},
			Tie,
		},
		{
			"paper vs scissors",
			RockPaperScissorsLizardSpock{},
			&Paper{},
			&Scissors{},
			Lose,
		},
		{
			"paper vs lizard",
			RockPaperScissorsLizardSpock{},
			&Paper{},
			&Lizard{},
			Lose,
		},
		{
			"paper vs spock",
			RockPaperScissorsLizardSpock{},
			&Paper{},
			&Spock{},
			Win,
		},
		{
			"scissors vs rock",
			RockPaperScissorsLizardSpock{},
			&Scissors{},
			&Rock{},
			Lose,
		},
		{
			"scissors vs paper",
			RockPaperScissorsLizardSpock{},
			&Scissors{},
			&Paper{},
			Win,
		},
		{
			"scissors vs scissors",
			RockPaperScissorsLizardSpock{},
			&Scissors{},
			&Scissors{},
			Tie,
		},
		{
			"scissors vs lizard",
			RockPaperScissorsLizardSpock{},
			&Scissors{},
			&Lizard{},
			Win,
		},
		{
			"scissors vs spock",
			RockPaperScissorsLizardSpock{},
			&Scissors{},
			&Spock{},
			Lose,
		},
		{
			"lizard vs rock",
			RockPaperScissorsLizardSpock{},
			&Lizard{},
			&Rock{},
			Lose,
		},
		{
			"lizard vs paper",
			RockPaperScissorsLizardSpock{},
			&Lizard{},
			&Paper{},
			Win,
		},
		{
			"lizard vs scissors",
			RockPaperScissorsLizardSpock{},
			&Lizard{},
			&Scissors{},
			Lose,
		},
		{
			"lizard vs lizard",
			RockPaperScissorsLizardSpock{},
			&Lizard{},
			&Lizard{},
			Tie,
		},
		{
			"lizard vs spock",
			RockPaperScissorsLizardSpock{},
			&Lizard{},
			&Spock{},
			Win,
		},
		{
			"spock vs rock",
			RockPaperScissorsLizardSpock{},
			&Spock{},
			&Rock{},
			Win,
		},
		{
			"spock vs paper",
			RockPaperScissorsLizardSpock{},
			&Spock{},
			&Paper{},
			Lose,
		},
		{
			"spock vs scissors",
			RockPaperScissorsLizardSpock{},
			&Spock{},
			&Scissors{},
			Win,
		},
		{
			"spock vs lizard",
			RockPaperScissorsLizardSpock{},
			&Spock{},
			&Lizard{},
			Lose,
		},
		{
			"spock vs spock",
			RockPaperScissorsLizardSpock{},
			&Spock{},
			&Spock{},
			Tie,
		},
	}
	for _, tt := range tests {
		actual := tt.rpsls.Result(tt.c1, tt.c2)
		if actual != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, actual)
		}
	}
}
