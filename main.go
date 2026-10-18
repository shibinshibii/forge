
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide command.")
		return
	}

	command := os.Args[1]

	if command != "create" {
		fmt.Println("Available command: create <project-name>")
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide project name after create.")
		return
	}

	projectName := os.Args[2]
	cleanProjectName := filepath.Base(projectName)

	if cleanProjectName != projectName {
		fmt.Println(
			"Invalid project name:",
			projectName,
			"using",
			cleanProjectName,
			"instead.",
		)
	}

	if dirExists(cleanProjectName) {
		fmt.Printf("The directory %s already exists\n", cleanProjectName)
		os.Exit(1)
	}

	// Pass the cleaned project name to filecontent.go.
	files := generateFiles(cleanProjectName)

	for _, value := range dirs {
		path := filepath.Join(cleanProjectName, value)

		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	for _, value := range files {
		filePath := filepath.Join(
			cleanProjectName,
			value.Folder,
			value.Name,
		)

		if err := os.WriteFile(filePath, value.Data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file %s: %v\n", value.Name, err)
			os.Exit(1)
		}
	}

	fmt.Printf("Project %s created successfully\n", cleanProjectName)
}
