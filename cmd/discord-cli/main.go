package main

import (
	"fmt"
	"os"

	"github.com/yourusername/discord-cli/internal/config"
	"github.com/yourusername/discord-cli/internal/discord"
	"github.com/yourusername/discord-cli/internal/ui"
)

func main() {
	// Initialize styles and UI components
	styles := ui.InitializeStyles()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error loading config:", err))
		os.Exit(1)
	}

	// Initialize Discord client
	client := discord.NewClient(cfg.Token)

	// Validate token if not already saved
	if cfg.Token == "" {
		token, err := ui.PromptForToken()
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error getting token:", err))
			os.Exit(1)
		}

		if err := client.ValidateToken(token); err != nil {
			fmt.Println(styles.ErrorStyle.Render("Invalid token:", err))
			os.Exit(1)
		}

		cfg.Token = token
		if err := config.Save(cfg); err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error saving config:", err))
			os.Exit(1)
		}
	}

	// Start the CLI
	if err := ui.StartCLI(client); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error running CLI:", err))
		os.Exit(1)
	}
}
