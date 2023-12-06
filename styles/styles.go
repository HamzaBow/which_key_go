package styles

import "github.com/charmbracelet/lipgloss"

var KeyStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var PrefixStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#2680d9"))

var GrayStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#777777"))

var CommandStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#d9d9268"))

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#FF0000")).
	PaddingLeft(1).
	PaddingRight(1)
