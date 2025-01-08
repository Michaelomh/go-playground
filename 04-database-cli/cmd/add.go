package cmd

import (
	"database-cli/db"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user to the database",
	Long:  `Add a new user to the database. The first argument is the name of the user.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAddUser,
}

func runAddUser(cmd *cobra.Command, args []string) error {
	// the args is the name of the employee
	name := args[0]
	fmt.Println("name:", name)
	user := db.User{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := db.Get().Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	fmt.Printf("Added new User(%s) to the database", args[0])
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
