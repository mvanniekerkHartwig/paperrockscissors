package game

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/davidspek/paperrockscissors/internal/engine"
	"github.com/davidspek/paperrockscissors/pkg"
)

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7"))
)

type Game struct {
	choices        []engine.Engine // the list of different game engines to choose from
	cursor         int             // the current cursor position for the engine selection
	scoreboard     *pkg.Scoreboard
	computerPlayer pkg.Player
	humanPlayer    pkg.Player
	isNewGame      bool
	nameInput      textinput.Model
	gameOver       bool
	playerResult   string
	engine         engine.Engine
}

func (g *Game) addPlayer(p *pkg.HumanPlayer) {
	g.humanPlayer = p
}

func New() *Game {
	return &Game{
		choices: []engine.Engine{
			engine.RockPaperScissors{},
			engine.RockPaperScissorsLizardSpock{},
		},
		computerPlayer: &pkg.ComputerPlayer{Name: "Computer"},
		isNewGame:      true,
		gameOver:       false,
		nameInput:      newNameInput(),
		scoreboard:     pkg.NewScoreboard(),
	}
}

func newNameInput() textinput.Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Placeholder = "Enter your name"
	return ti
}

func (g *Game) Init() tea.Cmd {
	return textinput.Blink
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Always quit if we receive a Ctrl+C or "q" key press
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return g, tea.Quit
		}
	}

	if g.isNewGame {
		cmd = g.handleNewGameUpdate(msg)
		return g, cmd
	}

	if g.engine == nil {
		cmd = g.handleEngineUpdate(msg)
		return g, cmd
	}

	if !g.gameOver {
		cmd = g.handleGameOverUpdate(msg)
		return g, cmd
	} else {
		// Allow the user to play again
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			// These keys should exit the program.
			case "ctrl+c", "q", "n":
				// used so the table is rendered before quitting
				g.gameOver = false
				return g, tea.Quit

			case "y":
				g.gameOver = false
			}
		}
	}
	return g, g.scoreboard.UpdateScores(msg, []pkg.Player{g.humanPlayer, g.computerPlayer})
}

// handleEngineUpdate handles the user's input when selecting a game engine
func (g *Game) handleEngineUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			g.engine = g.choices[g.cursor]
		case tea.KeyUp:
			if g.cursor > 0 {
				g.cursor--
			}
		case tea.KeyDown:
			if g.cursor < len(g.choices)-1 {
				g.cursor++
			}
		}
	}
	return g.scoreboard.UpdateScores(msg, []pkg.Player{g.humanPlayer, g.computerPlayer})
}

// handleNewGameUpdate handles the user's input when entering their name
func (g *Game) handleNewGameUpdate(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			g.addPlayer(&pkg.HumanPlayer{Name: g.nameInput.Value()})
			g.isNewGame = false
		}
	}
	g.nameInput, cmd = g.nameInput.Update(msg)
	return cmd
}

// handleGameOverUpdate handles the user's input during the game and updates the game state accordingly
func (g *Game) handleGameOverUpdate(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	if g.computerPlayer.Chosen() == nil {
		cmds = append(cmds, g.computerPlayer.SetChoice(msg, g.engine))
	}

	if g.humanPlayer.Chosen() == nil {
		cmds = append(cmds, g.humanPlayer.SetChoice(msg, g.engine))
	}

	if g.humanPlayer.Chosen() != nil && g.computerPlayer.Chosen() != nil {
		result := g.engine.Result(g.humanPlayer.Chosen(), g.computerPlayer.Chosen())
		g.humanPlayer.UpdateStats(result)
		g.computerPlayer.UpdateStats(result.Opposite())
		g.playerResult = result.String()

		g.humanPlayer.ClearChoice()
		g.computerPlayer.ClearChoice()
		g.gameOver = true
	}
	cmds = append(cmds, g.scoreboard.UpdateScores(msg, []pkg.Player{g.humanPlayer, g.computerPlayer}))
	return tea.Batch(cmds...)
}

func (g *Game) View() string {

	if g.isNewGame {
		return g.handleNewGameView()
	}

	if g.gameOver {
		return baseStyle.Render(fmt.Sprintf("You %s!\nWould you like to play again?", g.playerResult)) + g.helpLine()
	}

	if g.engine == nil {
		return g.handleEngineView()
	}

	s := "Choices\n\n"
	for _, choice := range g.engine.Choices() {
		// Render the available choices
		s += fmt.Sprintf("%s\n", choice.View())
	}

	var view string

	view += lipgloss.JoinHorizontal(lipgloss.Top, baseStyle.Render(s), baseStyle.Render(g.scoreboard.View()))

	return view + g.helpLine()
}

func (g *Game) handleNewGameView() string {
	return fmt.Sprintf(
		"What’s your name?\n\n%s\n",
		g.nameInput.View(),
	) + "\n"
}

func (g *Game) handleEngineView() string {
	s := "Choose a game\n\n"
	for i, choice := range g.choices {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if g.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name())
	}

	return baseStyle.Render(s) + g.helpLine()
}

func (g *Game) helpLine() string {
	helpLine := "\n"
	if !g.gameOver && g.engine != nil {

		for _, choice := range g.engine.Choices() {
			helpLine += fmt.Sprintf("%s: %s • ", choice.Key(), choice.View())
		}
	} else if g.gameOver {
		helpLine += "y: yes • n: no • "
	} else if g.engine == nil {
		helpLine += "enter: select • up/down: navigate • "
	}

	helpLine += "q: quit"
	return helpStyle.Render(helpLine)
}
