package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [taskID]",
	Short: "Complete a task",
	Long:  `Complete a task from the list by marking it as done.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCompleteTask,
}

func runCompleteTask(cmd *cobra.Command, args []string) error {
	taskID := args[0]

	records, err := ReadCSV()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}

	updated, found := markTaskComplete(records, taskID)
	if !found {
		return fmt.Errorf("no task found with ID: %s", taskID)
	}

	if err := WriteCSV(updated); err != nil {
		return fmt.Errorf("failed to update tasks: %w", err)
	}

	fmt.Printf("Task ID: %s is completed.\n", taskID)
	return nil
}

func markTaskComplete(records [][]string, taskID string) ([][]string, bool) {
	for i, record := range records {
		if record[IDColumn] == taskID {
			records[i][CompletedColumn] = "true"
			return records, true
		}
	}
	return records, false
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
