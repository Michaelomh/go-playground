package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the list.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAddTask,
}

func runAddTask(cmd *cobra.Command, args []string) error {
	taskName := args[0]

	records, err := ReadCSV()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}
	lastId, err := strconv.Atoi(records[len(records)-1][0])
	if err != nil {
		lastId = 0
	}

	records = append(records, addNewTask(lastId+1, taskName))
	if err := WriteCSV(records); err != nil {
		return fmt.Errorf("failed to add tasks: %w", err)
	}

	fmt.Printf("Task ID: %d is added.\n", lastId+1)
	return nil
}

func addNewTask(id int, taskName string) []string {
	return []string{strconv.Itoa(id), taskName, time.Now().Format("2006-01-01T15:04:05"), "false"}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
