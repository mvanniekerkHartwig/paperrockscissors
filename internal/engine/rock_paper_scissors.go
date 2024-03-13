package engine

type RockPaperScissors struct{}

func (rps RockPaperScissors) Name() string {
	return "Rock, Paper, Scissors"
}

func (rps RockPaperScissors) Choices() []Choice {
	return []Choice{
		&Rock{},
		&Paper{},
		&Scissors{},
	}
}

func (rps RockPaperScissors) Result(c1, c2 Choice) Result {
	if c1 == c2 {
		return Tie
	}
	if c1.Beats(c2) {
		return Win
	}
	return Lose
}

type Rock struct{}

func (r *Rock) Beats(c Choice) bool {
	switch c.(type) {
	case *Scissors:
		return true
	case *Lizard:
		return true
	}
	return false
}

func (r *Rock) View() string {
	return "Rock"
}

func (r *Rock) Key() string {
	return "r"
}

type Paper struct{}

func (p *Paper) Beats(c Choice) bool {
	switch c.(type) {
	case *Rock:
		return true
	case *Spock:
		return true
	}
	return false
}

func (p *Paper) View() string {
	return "Paper"
}

func (p *Paper) Key() string {
	return "p"
}

type Scissors struct{}

func (s *Scissors) Beats(c Choice) bool {
	switch c.(type) {
	case *Paper:
		return true
	case *Lizard:
		return true
	}
	return false
}

func (s *Scissors) View() string {
	return "Scissors"
}

func (s *Scissors) Key() string {
	return "s"
}
