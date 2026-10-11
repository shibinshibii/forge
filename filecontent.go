package main

var gitignoreContent = `# ##############################################################################
# Go Build Outputs & Binaries
# ##############################################################################
# Compiled binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Specific local build outputs (if you build to a local bin/ folder)
bin/
dist/

# ##############################################################################
# Testing and Code Coverage
# ##############################################################################
# Test binaries built automatically via
*.test

# Code coverage profiles and test artifacts
*.out
*.prof
*.coverprofile
coverage.txt

# ##############################################################################
# Dependencies and Workspace
# ##############################################################################
# Go workspace files (used for local multi-module development)
go.work
go.work.sum

# Dependency directories
# Un-comment the line below if you explicitly choose to vendor your dependencies.
# Otherwise, Go modules are cached globally, and the vendor directory shouldn't be committed.
# vendor/

# ##############################################################################
# Secrets, Environment Variables & Configuration
# ##############################################################################
.env
.env.local
*.local.toml

# ##############################################################################
# Operating System & IDE Files (Optional but Recommended)
# ##############################################################################
# macOS
.DS_Store

# Windows
Thumbs.db

# JetBrains (GoLand, IntelliJ)
.idea/
*.iml

# Visual Studio Code
.vscode/`

var gomodContent = `module github.com/yourusername/yourprojectname

go 1.27.1
`
var readmeContent = `# Your Project Name

This is a Go project generated using the Forge tool.
`