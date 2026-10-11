package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Name   string
	Folder string
	Data   []byte
}

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
	if command == "create" {
		if len(os.Args) < 3 {
			fmt.Println("Please provide project name after create.")
			return
		}
		projectName := os.Args[2]
		if dirExists(projectName) {
			fmt.Printf("The directory %s already exists", projectName)

		} else {
			for _, value := range dirs {
				path := filepath.Join(projectName, value)
				if err := os.MkdirAll(path, 0755); err != nil {
					log.Fatal("Something gone wrong", err)
				}
			}
			for _, value := range files {
				filePath := filepath.Join(projectName, value.Folder, value.Name)
				data := value.Data
				err := os.WriteFile(filePath, data, 0644)
				if err != nil {
					fmt.Println("Error Creating file", value.Name, err)
				}

			}

			fmt.Printf("Project %s created successfully", projectName)
		}
	} else {
		fmt.Println("Available command right now: create <project-name>")
	}
}
