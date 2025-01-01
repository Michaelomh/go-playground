package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks from the list.`,
	RunE:  listRecords,
}

func listRecords(cmd *cobra.Command, args []string) error {
	shouldShowAll, _ := cmd.Flags().GetBool("all")

	records, err := ReadCSV()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}

	listAllTasks(records, shouldShowAll)
	return nil
}

func listAllTasks(records [][]string, shouldShowAll bool) {
	w := new(tabwriter.Writer)
	defer w.Flush()
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)

	for _, record := range records {
		if !shouldShowAll && record[CompletedColumn] == "true" {
			continue // filter out completed tasks
		}
		fmt.Fprintln(w, strings.Join(record, "\t"))
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks regardless of completion status")
}
