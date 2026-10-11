package main

import (
	"fmt"
	"log"
	"os"
)
func main() {
	if len(os.Args) < 2{
		fmt.Println("Please provide command.")
		return
	}
	command := os.Args[1]
	if command == "create"{
		if len(os.Args) < 3 {
			fmt.Println("Please provide project name after create.")
			return
		}
		projectName := os.Args[2]
		err := os.Mkdir(projectName,0755)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("Project %s created successfully.", projectName)
	} else {

		fmt.Println("Available command right now: create <project-name>")
	}
}
