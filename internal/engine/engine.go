package engine

type Engine interface {
	// Returns the name of the game
	Name() string
	// Returns the choices available in the game
	Choices() []Choice
	// Returns the result of a match between two choices from the perspective of the first choice
	Result(Choice, Choice) Result
}

type Choice interface {
	// Returns true if the choice beats the argument
	Beats(Choice) bool
	// Returns a string representation of the choice
	View() string
	// Returns a string keybinding for the choice
	Key() string
}

type Result string

const (
	Win  Result = "win"
	Lose Result = "lose"
	Tie  Result = "tie"
)

func (r Result) String() string {
	return string(r)
}

func (r Result) Opposite() Result {
	switch r {
	case Win:
		return Lose
	case Lose:
		return Win
	}
	return Tie
}
