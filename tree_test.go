package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintTree(t *testing.T) {
	projectName := "test-project"
	wantTreePreview := `test-project/
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
`

	// Capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	oldStdout := os.Stdout
	os.Stdout = w

	printTree(projectName)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}

	got := buf.String()
	if got != wantTreePreview {
		t.Errorf("printTree(%q) output mismatch:\ngot:\n%s\nwant:\n%s", projectName, got, wantTreePreview)
	}
}
