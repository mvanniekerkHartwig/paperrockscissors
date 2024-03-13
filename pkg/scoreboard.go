package pkg

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Scoreboard struct {
	table.Model
}

func NewScoreboard() *Scoreboard {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "Name", Width: 10},
		{Title: "Wins", Width: 10},
		{Title: "Losses", Width: 10},
		{Title: "Ties", Width: 10},
	}

	rows := []table.Row{}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(4),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	// Set the cursor to -1 so that it doesn't show up
	t.SetCursor(-1)

	return &Scoreboard{
		t,
	}
}

func (s *Scoreboard) UpdateScores(msg tea.Msg, players []Player) (cmd tea.Cmd) {
	rows := []table.Row{}
	for _, p := range players {
		rows = append(rows, p.StatsRow())
	}
	sorted := sortRows(rows)
	for i, r := range sorted {
		sorted[i] = append([]string{lipgloss.NewStyle().Width(1).Render(fmt.Sprintf("%d", i+1))}, r...)
	}
	s.SetRows(sorted)
	s.Model, cmd = s.Update(msg)
	return cmd
}

func sortRows(rows []table.Row) []table.Row {
	sort.Slice(rows, func(i, j int) bool {
		return rows[i][1] > rows[j][1]
	})
	return rows
}
