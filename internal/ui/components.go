package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/matu6968/project-lite-go/internal/discord"
)

// PromptForToken prompts the user for their Discord token
func PromptForToken() (string, error) {
	var token string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your Discord token").
				Value(&token).
				Password(true),
		),
	)

	err := form.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}

	return token, nil
}

// StartCLI starts the main CLI interface
func StartCLI(client *discord.Client) error {
	fmt.Println("project-lite-go")
	fmt.Println("A CLI Discord client")
	fmt.Println("Type 'help' for available commands")

	// TODO: Implement the main CLI loop
	// This is a placeholder that just returns nil
	return nil
}

// RenderError formats and renders an error message
func RenderError(err error) string {
	return fmt.Sprintf("Error: %v", err)
}
