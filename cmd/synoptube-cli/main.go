package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command for the application.
var rootCmd = &cobra.Command{
	Use:   "syt1",
	Short: "A simple command-line application.",
	Long: `This is a simple CLI application that uses the Cobra library to handle
command-line arguments and flags. It provides a single command to summarize a YouTube video.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This will be executed if no subcommand is given.
		cmd.Help()
	},
}

// summarizeCmd represents the subcommand to summarize a YouTube video.
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarizes a YouTube video from its URL.",
	Long:  `The 'summarize' command takes a YouTube URL as a flag and returns a summary.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the value of the 'url' flag.
		youtubeURL, _ := cmd.Flags().GetString("url")

		// Check if the required 'url' flag was provided.
		if youtubeURL == "" {
			fmt.Println("Error: The --url flag is required.")
			cmd.Help()
			os.Exit(1)
		}

		// Print a message acknowledging the input.
		fmt.Printf("Received YouTube URL: %s\n", youtubeURL)
		fmt.Println("Processing for summarization...")

		// A placeholder for the actual summarization logic.
		// In a real application, you would pass this URL to your core logic.
		log.Println("CLI application finished.")
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is the main entry point for the CLI.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// main is the entry point of the application.
func main() {
	Execute()
}
