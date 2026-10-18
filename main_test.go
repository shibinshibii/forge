package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var forgeBinPath string

func TestMain(m *testing.M) {
	tempDir, err := os.MkdirTemp("", "forge-test-bin-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create temp dir for forge test binary: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	forgeBinPath = filepath.Join(tempDir, "forge")
	buildCmd := exec.Command("go", "build", "-o", forgeBinPath, ".")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to build forge binary: %v\n%s\n", err, out)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestDirExists(t *testing.T) {
	tempDir := t.TempDir()

	// 1. Existing directory
	subDir := filepath.Join(tempDir, "existing-dir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}
	if !dirExists(subDir) {
		t.Errorf("dirExists(%q) = false, want true for an existing directory", subDir)
	}

	// 2. Non-existent path
	nonExistent := filepath.Join(tempDir, "does-not-exist")
	if dirExists(nonExistent) {
		t.Errorf("dirExists(%q) = true, want false for non-existent path", nonExistent)
	}

	// 3. Existing file (not a directory)
	filePath := filepath.Join(tempDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	if dirExists(filePath) {
		t.Errorf("dirExists(%q) = true, want false for a regular file", filePath)
	}

	// 4. Empty string
	if dirExists("") {
		t.Errorf("dirExists(\"\") = true, want false for empty path")
	}
}

func TestCLI_NoArguments(t *testing.T) {
	cmd := exec.Command(forgeBinPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected exit code 0 when no args provided, got error: %v", err)
	}

	output := strings.TrimSpace(string(out))
	expected := "Please provide command."
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}

func TestCLI_InvalidCommand(t *testing.T) {
	cmd := exec.Command(forgeBinPath, "unknown")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected exit code 0 on invalid command, got error: %v", err)
	}

	output := strings.TrimSpace(string(out))
	expected := "Available command: create <project-name>"
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}

func TestCLI_MissingProjectName(t *testing.T) {
	cmd := exec.Command(forgeBinPath, "create")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected exit code 0 when missing project name, got error: %v", err)
	}

	output := strings.TrimSpace(string(out))
	expected := "Please provide project name after create."
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}

func TestCLI_CreateSuccess(t *testing.T) {
	workDir := t.TempDir()
	projectName := "my-service"

	cmd := exec.Command(forgeBinPath, "create", projectName)
	cmd.Dir = workDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("forge create failed: %v\noutput: %s", err, out)
	}

	output := strings.TrimSpace(string(out))
	expectedOutput := "Project my-service created successfully"
	if output != expectedOutput {
		t.Errorf("got output %q, want %q", output, expectedOutput)
	}

	projectDir := filepath.Join(workDir, projectName)
	if !dirExists(projectDir) {
		t.Fatalf("expected project directory %q to exist", projectDir)
	}

	// Verify all directories exist
	expectedDirs := []string{
		"cmd/server",
		"internal/config",
		"internal/handlers",
		"migrations",
	}
	for _, relDir := range expectedDirs {
		targetDir := filepath.Join(projectDir, relDir)
		if !dirExists(targetDir) {
			t.Errorf("expected directory %q does not exist", targetDir)
		}
	}

	// Verify all files exist
	expectedFiles := []string{
		"README.md",
		".gitignore",
		"go.mod",
		"cmd/server/main.go",
		"internal/config/config.go",
		"internal/handlers/handler.go",
	}
	for _, relFile := range expectedFiles {
		targetFile := filepath.Join(projectDir, relFile)
		info, err := os.Stat(targetFile)
		if err != nil {
			t.Errorf("expected file %q does not exist: %v", targetFile, err)
			continue
		}
		if info.IsDir() {
			t.Errorf("expected %q to be a file, but got directory", targetFile)
		}
		if info.Size() == 0 {
			t.Errorf("expected %q to be non-empty", targetFile)
		}
	}

	// Verify go.mod content
	goModBytes, err := os.ReadFile(filepath.Join(projectDir, "go.mod"))
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}
	if !strings.Contains(string(goModBytes), "module example.com/my-service") {
		t.Errorf("go.mod does not contain expected module path: %s", string(goModBytes))
	}

	// Verify that the generated project compiles with `go vet`
	vetCmd := exec.Command("go", "vet", "./...")
	vetCmd.Dir = projectDir
	if vetOut, err := vetCmd.CombinedOutput(); err != nil {
		t.Errorf("go vet failed on generated project: %v\noutput: %s", err, vetOut)
	}

	// Verify that the generated project builds successfully with `go build ./cmd/server`
	buildCmd := exec.Command("go", "build", "./cmd/server")
	buildCmd.Dir = projectDir
	if buildOut, err := buildCmd.CombinedOutput(); err != nil {
		t.Errorf("go build failed on generated server: %v\noutput: %s", err, buildOut)
	}
}

func TestCLI_CreateDirectoryAlreadyExists(t *testing.T) {
	workDir := t.TempDir()
	projectName := "existing-app"

	// Pre-create the directory
	if err := os.Mkdir(filepath.Join(workDir, projectName), 0755); err != nil {
		t.Fatalf("failed to pre-create directory: %v", err)
	}

	cmd := exec.Command(forgeBinPath, "create", projectName)
	cmd.Dir = workDir
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected command to fail when directory already exists, but it succeeded")
	}

	output := strings.TrimSpace(string(out))
	expectedOutput := "The directory existing-app already exists"
	if output != expectedOutput {
		t.Errorf("got %q, want %q", output, expectedOutput)
	}
}

func TestCLI_CreateSanitizesNestedPath(t *testing.T) {
	workDir := t.TempDir()
	rawProjectName := "nested/path/sample-api"
	cleanName := "sample-api"

	cmd := exec.Command(forgeBinPath, "create", rawProjectName)
	cmd.Dir = workDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("forge create failed: %v\noutput: %s", err, out)
	}

	output := string(out)
	expectedWarning := fmt.Sprintf("Invalid project name: %s using %s instead.", rawProjectName, cleanName)
	if !strings.Contains(output, expectedWarning) {
		t.Errorf("output missing warning %q; got: %s", expectedWarning, output)
	}

	expectedSuccess := fmt.Sprintf("Project %s created successfully", cleanName)
	if !strings.Contains(output, expectedSuccess) {
		t.Errorf("output missing success %q; got: %s", expectedSuccess, output)
	}

	// Check that the project directory was created as cleanName in workDir
	projectDir := filepath.Join(workDir, cleanName)
	if !dirExists(projectDir) {
		t.Errorf("expected directory %q to exist", projectDir)
	}
}
