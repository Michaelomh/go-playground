package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "database-cli",
	Short: "A quick application that interfaces with the database via cli",
	Long: `A CLI application which is able to interact with a database via cli. 
	
	This application acts like a HR employee management system.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
