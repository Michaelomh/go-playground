package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task from the list.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDeleteTask,
}

func runDeleteTask(cmd *cobra.Command, args []string) error {
	taskID := args[0]

	records, err := ReadCSV()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}

	updated, found := removeRecords(records, taskID)
	if !found {
		return fmt.Errorf("no task found with ID: %s. Please try again.", taskID)
	}

	if err := WriteCSV(updated); err != nil {
		return fmt.Errorf("failed to update tasks: %w", err)
	}

	fmt.Printf("Task ID: %s is completed.\n", taskID)
	return nil
}

func removeRecords(records [][]string, taskID string) ([][]string, bool) {
	foundId := false
	for i, record := range records {
		if record[IDColumn] == taskID {
			records = append(records[:i], records[i+1:]...)
			foundId = true
		}
	}
	return records, foundId
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
