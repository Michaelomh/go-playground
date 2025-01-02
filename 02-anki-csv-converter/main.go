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

// func main() {
// 	form := huh.NewForm(
// 		huh.NewGroup(
// 			huh.NewInput().
// 				Title("What is the directory of anki file you wish to convert?").
// 				Value(&directory).
// 				Placeholder(fmt.Sprintf("Default: %s", cmd.DefaultDirectory)).
// 				Validate(func(directory string) error {
// 					directory = cmd.GetDefaultDirectory(directory)
// 					if !cmd.IsValidDirectory(directory) {
// 						return errors.New("Directory is not a valid directory. Please try again")
// 					}
// 					return nil
// 				}),
// 		),

// 		huh.NewGroup(
// 			huh.NewSelect[string]().
// 				Title("What is the file name?").
// 				OptionsFunc(func() []huh.Option[string] {
// 					directory = cmd.GetDefaultDirectory(directory)
// 					opts := getAllFilesInDirectory(directory)
// 					return huh.NewOptions(opts...)
// 				}, &directory).
// 				Value(&file),
// 		),

// 		huh.NewGroup(
// 			huh.NewInput().
// 				Title("What would be the output file name?").
// 				Value(&outputFile).
// 				Placeholder(fmt.Sprintf("Default: %s", cmd.DefaultOutputFile)).
// 				Validate(func(outputFile string) error {
// 					outputFile = cmd.GetDefaultOutputFile(outputFile)
// 					if !cmd.IsValidTextFile(outputFile) {
// 						return errors.New("Output is not a valid text file. Please try again")
// 					}
// 					return nil
// 				}),
// 		),
// 	)

// 	err := form.WithTheme(huh.ThemeDracula()).Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	directory = cmd.GetDefaultDirectory(directory)
// 	outputFile = cmd.GetDefaultOutputFile(outputFile)

// 	records, err := cmd.ReadTextFile(directory + "/" + file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Writing to:", directory+"/"+outputFile)
// 	cmd.WriteCSV(records, directory+"/"+outputFile)
// 	fmt.Println("Writing completed")
// }

// func getAllFilesInDirectory(dirPath string) []string {
// 	txtFiles := []string{}
// 	files, err := os.ReadDir(dirPath)
// 	if err != nil {
// 		log.Fatal("Error reading directory:", err)
// 		return txtFiles
// 	}

// 	for _, file := range files {
// 		if strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1] == "txt" {
// 			txtFiles = append(txtFiles, file.Name())
// 		}
// 	}

// 	return txtFiles
// }
