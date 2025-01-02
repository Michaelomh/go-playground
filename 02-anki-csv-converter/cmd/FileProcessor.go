package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
)

// FormData holds the user input data from the form
type FormData struct {
	directory  string
	file       string
	outputFile string
}

// FileProcessor handles file processing operations
type FileProcessor struct {
	formData FormData
}

func NewFileProcessor() *FileProcessor {
	return &FileProcessor{}
}

// getTextFiles returns all .txt files in the given directory
func (fp *FileProcessor) GetTextFiles(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var txtFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			txtFiles = append(txtFiles, file.Name())
		}
	}
	return txtFiles, nil
}

// createForm builds and returns the form UI
func (fp *FileProcessor) CreateForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the directory of anki file you wish to convert?").
				Value(&fp.formData.directory).
				Placeholder(fmt.Sprintf("Default: %s", DefaultDirectory)).
				Validate(fp.ValidateDirectory),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is the file name?").
				OptionsFunc(func() []huh.Option[string] {
					directory := GetDefaultDirectory(fp.formData.directory)
					files, err := fp.GetTextFiles(directory)
					if err != nil {
						log.Fatalf("Error getting text files: %v", err)
						return nil
					}
					return huh.NewOptions(files...)
				}, &fp.formData.directory).
				Value(&fp.formData.file),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What would be the output file name?").
				Value(&fp.formData.outputFile).
				Placeholder(fmt.Sprintf("Default: %s", DefaultOutputFile)).
				Validate(fp.ValidateOutputFile),
		),
	)
}

// processFiles handles the file conversion process
func (fp *FileProcessor) ProcessFiles() error {
	directory := GetDefaultDirectory(fp.formData.directory)
	outputFile := GetDefaultOutputFile(fp.formData.outputFile)

	inputPath := filepath.Join(directory, fp.formData.file)
	outputPath := filepath.Join(directory, outputFile)

	records, err := ReadTextFile(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	fmt.Printf("Writing to: %s\n", outputPath)
	if err := WriteCSV(records, outputPath); err != nil {
		return fmt.Errorf("error writing CSV: %w", err)
	}

	fmt.Println("Writing completed.🥳")
	return nil
}
