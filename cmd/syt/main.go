package main

import (
	"fmt"
	"os"
)

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
