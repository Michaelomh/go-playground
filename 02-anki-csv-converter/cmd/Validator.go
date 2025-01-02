package cmd

import (
	"fmt"
	"os"
	"strings"
)

// validateDirectory checks if the directory is valid
func (fp *FileProcessor) ValidateDirectory(directory string) error {
	directory = fp.GetDefaultDirectory(directory)
	if !IsValidDirectory(directory) {
		return fmt.Errorf("directory is not valid, please try again")
	}
	return nil
}

// validateOutputFile checks if the output file is valid
func (fp *FileProcessor) ValidateOutputFile(outputFile string) error {
	outputFile = fp.GetDefaultOutputFile(outputFile)
	if !IsValidTextFile(outputFile) {
		return fmt.Errorf("output is not a valid text file, please try again")
	}
	return nil
}

func IsValidDirectory(p string) bool {
	fileInfo, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsValidTextFile(f string) bool {
	splitString := strings.Split(f, ".")
	return splitString[len(splitString)-1] == "txt"
}
