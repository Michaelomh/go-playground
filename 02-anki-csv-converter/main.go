package main

import (
	"anki-cc/cmd"
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	processor := cmd.NewFileProcessor()
	form := processor.CreateForm()

	if err := form.WithTheme(huh.ThemeDracula()).Run(); err != nil {
		log.Fatalf("Error running form: %v", err)
	}

	if err := processor.ProcessFiles(); err != nil {
		log.Fatalf("Error processing files: %v", err)
	}
}
