package data

import (
	"slices"
	"strings"
	"testing"
)

func TestDirs(t *testing.T) {
	expectedDirs := []string{
		"cmd/server",
		"internal/config",
		"internal/handlers",
		"migrations",
	}

	if len(Dirs) != len(expectedDirs) {
		t.Fatalf("got %d Dirs, want %d Dirs", len(Dirs), len(expectedDirs))
	}

	for _, expected := range expectedDirs {
		if !slices.Contains(Dirs, expected) {
			t.Errorf("expected directory %q not found in Dirs list", expected)
		}
	}

	// Verify no empty paths, no leading slashes, and no duplicate paths
	seen := make(map[string]bool)
	for _, dir := range Dirs {
		if strings.TrimSpace(dir) == "" {
			t.Errorf("Dirs contains an empty directory path")
		}
		if strings.HasPrefix(dir, "/") || strings.HasPrefix(dir, "\\") {
			t.Errorf("dir path %q should be relative, but has leading slash", dir)
		}
		if seen[dir] {
			t.Errorf("duplicate directory path found in Dirs: %q", dir)
		}
		seen[dir] = true
	}
}

func TestFileStruct(t *testing.T) {
	file := File{
		Name:   "example.txt",
		Folder: "cmd/example",
		Data:   []byte("sample data"),
	}

	if file.Name != "example.txt" {
		t.Errorf("got file.Name = %q, want %q", file.Name, "example.txt")
	}
	if file.Folder != "cmd/example" {
		t.Errorf("got file.Folder = %q, want %q", file.Folder, "cmd/example")
	}
	if string(file.Data) != "sample data" {
		t.Errorf("got file.Data = %q, want %q", string(file.Data), "sample data")
	}
}
