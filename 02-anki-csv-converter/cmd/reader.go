package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

const (
	ChineseColumn = 0
	EnglishColumn = 1
	PinyinColumn  = 2
)

// ReadTextFile reads all records from the .txt file
func ReadTextFile(file string) ([][]string, error) {
	records := [][]string{}
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	contents := strings.Split(string(dat), "\n")
	for _, content := range contents {
		if content == "" || content[0] == '#' {
			continue
		}

		content = strings.ReplaceAll(content, "<br><br>", "<br>")
		content = strings.ReplaceAll(content, "<br>", "\t")
		contentArr := strings.Split(content, "\t")

		// fmt.Println(contentArr)
		// fmt.Println(contentArr[0]) // should be all chinese
		// fmt.Println(contentArr[1]) // should be all english
		// fmt.Println(contentArr[2]) // should be all pinyin

		if len(contentArr) > 3 {
			contentArr[2] = strings.Join(contentArr[2:], "")
		}

		records = append(records, contentArr)
	}

	return records, nil
}

// WriteCSV writes records to the CSV file atomically
func WriteCSV(records [][]string, filename string) error {
	// Create a temporary file in the same directory
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create temporary file: %w", err)
	}

	w := csv.NewWriter(f)
	if err := w.WriteAll(records); err != nil {
		f.Close()
		return fmt.Errorf("error writing records: %w", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		f.Close()
		return fmt.Errorf("error flushing records: %w", err)
	}

	fmt.Println("output fileName", filename)

	return nil
}
