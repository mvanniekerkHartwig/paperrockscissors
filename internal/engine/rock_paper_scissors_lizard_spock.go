package engine

type RockPaperScissorsLizardSpock struct{}

func (rps RockPaperScissorsLizardSpock) Name() string {
	return "Rock, Paper, Scissors, Lizard, Spock"
}

func (rps RockPaperScissorsLizardSpock) Choices() []Choice {
	return []Choice{
		&Rock{},
		&Paper{},
		&Scissors{},
		&Lizard{},
		&Spock{},
	}
}

func (rps RockPaperScissorsLizardSpock) Result(c1, c2 Choice) Result {
	if c1 == c2 {
		return Tie
	}
	if c1.Beats(c2) {
		return Win
	}
	return Lose
}

type Lizard struct{}

func (l *Lizard) Beats(c Choice) bool {
	switch c.(type) {
	case *Spock:
		return true
	case *Paper:
		return true
	}
	return false
}

func (l *Lizard) View() string {
	return "Lizard"
}

func (l *Lizard) Key() string {
	return "l"
}

type Spock struct{}

func (s *Spock) Beats(c Choice) bool {
	switch c.(type) {
	case *Scissors:
		return true
	case *Rock:
		return true
	}
	return false
}

func (s *Spock) View() string {
	return "Spock"
}

func (s *Spock) Key() string {
	return "v"
}
