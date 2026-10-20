package main

import "fmt"

func printTree(projectName string) {
	fmt.Printf(`%s/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   └── handlers/
│       └── handler.go
├── migrations/
├── .gitignore
├── README.md
└── go.mod
`, projectName)
}
