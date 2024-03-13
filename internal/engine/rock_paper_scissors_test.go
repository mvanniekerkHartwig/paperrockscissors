package engine

import "testing"

func TestRPSChoices(t *testing.T) {
	rps := RockPaperScissors{}
	choices := rps.Choices()
	if len(choices) != 3 {
		t.Errorf("expected 3 choices, got %d", len(choices))
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
}

func TestRPSResult(t *testing.T) {
	tests := []struct {
		name     string
		rps      RockPaperScissors
		c1, c2   Choice
		expected Result
	}{
		{
			"rock vs rock",
			RockPaperScissors{},
			&Rock{},
			&Rock{},
			Tie,
		},
		{
			"rock vs paper",
			RockPaperScissors{},
			&Rock{},
			&Paper{},
			Lose,
		},
		{
			"rock vs scissors",
			RockPaperScissors{},
			&Rock{},
			&Scissors{},
			Win,
		},
		{
			"paper vs rock",
			RockPaperScissors{},
			&Paper{},
			&Rock{},
			Win,
		},
		{
			"paper vs paper",
			RockPaperScissors{},
			&Paper{},
			&Paper{},
			Tie,
		},
		{
			"paper vs scissors",
			RockPaperScissors{},
			&Paper{},
			&Scissors{},
			Lose,
		},
		{
			"scissors vs rock",
			RockPaperScissors{},
			&Scissors{},
			&Rock{},
			Lose,
		},
		{
			"scissors vs paper",
			RockPaperScissors{},
			&Scissors{},
			&Paper{},
			Win,
		},
		{
			"scissors vs scissors",
			RockPaperScissors{},
			&Scissors{},
			&Scissors{},
			Tie,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.rps.Result(tt.c1, tt.c2)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
