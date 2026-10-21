package tree

import "fmt"

func PrintTree(projectName string) {
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
