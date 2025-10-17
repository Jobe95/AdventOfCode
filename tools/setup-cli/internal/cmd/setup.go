package cmd

import (
	"fmt"

	"github.com/Jobe95/AdventOfCode/internal/config"
	"github.com/Jobe95/AdventOfCode/internal/setup"
	"github.com/Jobe95/AdventOfCode/internal/ui"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up a new Advent of Code day",
	Long:  `Interactive setup for a new Advent of Code puzzle day. Select year, day, and languages.`,
	RunE:  runSetup,
}

func runSetup(cmd *cobra.Command, args []string) error {
	session, err := config.GetSession()
	if err != nil {
		return fmt.Errorf("failed to check session: %w", err)
	}

	if session == "" {
		return fmt.Errorf("session not configured\n\nRun: aoc config set-session <cookie>\n\nGet your session cookie from https://adventofcode.com\n(DevTools > Application > Cookies > session)")
	}

	selections, err := ui.RunInteractive()
	if err != nil {
		return err
	}

	if err := setup.Execute(selections); err != nil {
		return err
	}

	return nil
}
