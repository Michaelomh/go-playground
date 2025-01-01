package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// ReadCSV reads all records from the CSV file
func ReadCSV() ([][]string, error) {
	f, err := os.OpenFile(CSVFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read records: %w", err)
	}

	return records, nil
}

// WriteCSV writes records to the CSV file atomically
func WriteCSV(records [][]string) error {
	// Create a temporary file in the same directory
	dir := filepath.Dir(CSVFileName)
	tmpFile, err := os.CreateTemp(dir, "*.csv")
	if err != nil {
		return fmt.Errorf("unable to create temporary file: %w", err)
	}
	tmpName := tmpFile.Name()
	defer os.Remove(tmpName) // Clean up the temp file in case of failure

	w := csv.NewWriter(tmpFile)
	if err := w.WriteAll(records); err != nil {
		tmpFile.Close()
		return fmt.Errorf("error writing records: %w", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		tmpFile.Close()
		return fmt.Errorf("error flushing records: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("error closing temporary file: %w", err)
	}

	// automatically replace the original file
	if err := os.Rename(tmpName, CSVFileName); err != nil {
		return fmt.Errorf("error replacing original file: %w", err)
	}

	return nil
}
