package cmd

import (
	"os"
	"strings"
)

const (
	DefaultDirectory  = "/Users/michaelomh/Downloads"
	DefaultOutputFile = "anki-ccc.txt"
)

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

func GetDefaultDirectory(d string) string {
	if strings.TrimSpace(d) == "" {
		d = DefaultDirectory
	}
	return d
}

func GetDefaultOutputFile(d string) string {
	if strings.TrimSpace(d) == "" {
		d = DefaultOutputFile
	}
	return d
}
