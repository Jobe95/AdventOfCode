package cmd

import (
	"fmt"

	"github.com/Jobe95/AdventOfCode/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage your Advent of Code configuration, including session cookie.`,
}

var setSessionCmd = &cobra.Command{
	Use:   "set-session [session-cookie]",
	Short: "Set your AOC session cookie",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.SetSession(args[0]); err != nil {
			return err
		}
		fmt.Println("✓ Session saved")
		return nil
	},
}

var showSessionCmd = &cobra.Command{
	Use:   "show-session",
	Short: "Show your current AOC session cookie",
	RunE: func(cmd *cobra.Command, args []string) error {
		session, err := config.GetSession()
		if err != nil {
			return err
		}
		if session == "" {
			fmt.Println("No session configured")
		} else {
			fmt.Println(session)
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(setSessionCmd)
	configCmd.AddCommand(showSessionCmd)
}
