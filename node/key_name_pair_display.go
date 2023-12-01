package node

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func (nd Node) Display() {
	columns := []table.Column{
		{Title: "Key", Width: 3},
		{Title: "Name", Width: 10},
	}

	rows := []table.Row{}
	for k, v := range nd.Children {
		kk := table.Row{k, v.Name}
		rows = append(rows, kk)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		BorderTop(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("250"))
	t.SetStyles(s)

	fmt.Println(t.View())
}
