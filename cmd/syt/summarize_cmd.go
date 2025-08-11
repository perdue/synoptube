package main

import (
	"fmt"
	"log"
	"os"

	"github.com/perdue/synoptube/pkg/summarize"
	"github.com/spf13/cobra"
)

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

		if err := summarize.ProcessVideo(); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// A placeholder for the actual summarization logic.
		// In a real application, you would pass this URL to your core logic.
		log.Println("CLI application finished.")
	},
}
