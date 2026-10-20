package main

import "fmt"

func printHelp() {
	fmt.Printf(`Forge - A tiny Go project scaffolding CLI

Generate a basic Go API project structure without a web framework.

USAGE
  forge <command> [arguments]

COMMANDS
  create <project_name>    Create a new Go API project
  tree [project_name]      Preview the generated project structure
  help                     Show this help message

EXAMPLES
  forge create my-api
  forge tree
  forge tree my-api
  forge help

ABOUT
  Forge is a lightweight CLI tool built with Go.
  It helps developers scaffold a basic Go API project
  using the Go standard library, without a web framework.

GitHub
  https://github.com/shibinshibii/forge
`)
}
