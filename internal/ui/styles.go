package ui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Styles struct {
	ErrorStyle   lipgloss.Style
	SuccessStyle lipgloss.Style
	InfoStyle    lipgloss.Style
	InputStyle   lipgloss.Style
}

func InitializeStyles() *Styles {
	// We'll use the output for color support detection
	_ = termenv.NewOutput(os.Stdout)

	return &Styles{
		ErrorStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true),

		SuccessStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true),

		InfoStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0000FF")),

		InputStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#333333")).
			Padding(0, 1),
	}
}
