package engine

import "testing"

func TestResultOpposite(t *testing.T) {
	tests := []struct {
		input    Result
		expected Result
	}{
		{
			Win,
			Lose,
		},
		{
			Lose,
			Win,
		},
		{
			Tie,
			Tie,
		},
	}
	for _, tt := range tests {
		actual := tt.input.Opposite()
		if actual != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, actual)
		}
	}
}
