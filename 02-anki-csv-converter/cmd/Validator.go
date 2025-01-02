package cmd

import "fmt"

// validateDirectory checks if the directory is valid
func (fp *FileProcessor) ValidateDirectory(directory string) error {
	directory = GetDefaultDirectory(directory)
	if !IsValidDirectory(directory) {
		return fmt.Errorf("directory is not valid, please try again")
	}
	return nil
}

// validateOutputFile checks if the output file is valid
func (fp *FileProcessor) ValidateOutputFile(outputFile string) error {
	outputFile = GetDefaultOutputFile(outputFile)
	if !IsValidTextFile(outputFile) {
		return fmt.Errorf("output is not a valid text file, please try again")
	}
	return nil
}
