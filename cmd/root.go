package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes",
	Short: "Notes CLI - fast note-taking for developers",
	Long:  "A terminal-based note-taking CLI with Git sync, tags, and optional encryption.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Notes CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
