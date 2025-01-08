/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database-cli/db"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List users data from the database",
	Long:  `List all users in the User table.`,
	RunE:  listUsers,
}

func listUsers(cmd *cobra.Command, args []string) error {
	var users []db.User
	db.Get().Find(&users)

	w := new(tabwriter.Writer)
	defer w.Flush()
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)

	for _, record := range users {
		print := []string{fmt.Sprintf("%d", record.ID), record.Name, record.CreatedAt.Format("2006-01-02"), record.UpdatedAt.Format("2006-01-02")}
		fmt.Fprintln(w, strings.Join(print, "\t"))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
