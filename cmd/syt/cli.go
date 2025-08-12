package main

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command for the application.
var rootCmd = &cobra.Command{
	Use: "syt",
	Long: `Synoptube Command Line Interface:
    "syt" (pronounced like "sight") is a CLI that facilitates
    building summaries of YouTube videos.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This will be executed if no subcommand is given.
		cmd.Help()
	},
}

// init function is where we define the subcommands and their flags.
// This is called before main().
func init() {
	// Add the 'summarize' subcommand to the root command.
	rootCmd.AddCommand(summarizeCmd)

	// Add a string flag to the 'summarize' subcommand.
	// We define a string flag named "url" with a short-hand 'u'.
	summarizeCmd.Flags().StringP("url", "u", "", "The YouTube URL to summarize.")
}
