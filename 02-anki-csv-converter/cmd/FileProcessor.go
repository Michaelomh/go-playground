package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

const (
	DefaultDirectory  = "/Users/michaelomh/Downloads"
	DefaultOutputFile = "anki-ccc.txt"
)

type FormData struct {
	directory  string
	file       string
	outputFile string
}

type FileProcessor struct {
	formData FormData
}

func NewFileProcessor() *FileProcessor {
	return &FileProcessor{}
}

// getTxtFilesInDir returns all .txt files in the given directory
func (fp *FileProcessor) getTxtFilesInDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
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

// createForm builds and returns the form TUI
// users would populate the data accordingly to the TUI
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
					directory := fp.GetDefaultDirectory(fp.formData.directory)
					files, err := fp.getTxtFilesInDir(directory)
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

// ProcessFiles handles the file conversion process
func (fp *FileProcessor) ProcessFiles() error {
	directory := fp.GetDefaultDirectory(fp.formData.directory)
	outputFile := fp.GetDefaultOutputFile(fp.formData.outputFile)

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

func (fp *FileProcessor) GetDefaultDirectory(d string) string {
	if strings.TrimSpace(d) == "" {
		d = DefaultDirectory
	}
	return d
}

func (fp *FileProcessor) GetDefaultOutputFile(d string) string {
	if strings.TrimSpace(d) == "" {
		d = DefaultOutputFile
	}
	return d
}
