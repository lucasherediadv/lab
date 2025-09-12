package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A simple example of a CLI app with Cobra",
	Long:  `This is a simple CLI application created to demonstrate how Cobra works.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is the default action when no subcommand is specified.
		fmt.Println("Welcome to cobra!")
		fmt.Println("Try running `cobra hello`")
	},
}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting",
	Long:  `This command prints a friendly greeting to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from cobra!")
	},
}

func main() {
	// Add the helloCmd as a subcommand to the rootCmd
	rootCmd.AddCommand(helloCmd)

	// Execute the root command. This is what starts the CLI app.
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
