# Forge

A small Go CLI that scaffolds a predefined Go API project.

## Goal

Make starting a basic Go API repeatable. V1 intentionally supports one Go template and uses Go's standard library.

## Prerequisite: Go

Forge is installed with `go install`, so Go must be installed first.

- [Download Go](https://go.dev/dl/)
- [Official installation guide and troubleshooting](https://go.dev/doc/install)

Verify the installation with `go version`. If the `go` command is not found, follow the official guide to configure your `PATH`.

## Install

```sh
go install github.com/shibinshibii/forge@v0.1.0
```

Make sure Go's binary directory is on your `PATH` so the `forge` command can be found.

## Usage

```sh
forge <command> [arguments]
```

### Commands

| Command | Description |
| --- | --- |
| `forge create <project_name>` | Create a new Go API project |
| `forge tree [project_name]` | Preview the generated project structure |
| `forge help` | Show help message and usage details |

### Create a project

```sh
forge create my-api
cd my-api
go run ./cmd/server
```

The server listens on port `8080` by default. Set `PORT` to use another port.

Health check:

```sh
curl http://localhost:8080/health
```

Returns `{"status":"ok"}`.

### Preview project structure

Preview the directory structure without generating files on disk:

```sh
# Preview with default placeholder ("your-project-name")
forge tree

# Preview with a specific project name
forge tree my-api
```

Output:

```text
my-api/
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
```

### Help

Display CLI usage, available commands, and examples:

```sh
forge help
```

## Develop

```sh
git clone https://github.com/shibinshibii/forge.git
cd forge
go test ./...
go run . create test-api
go run . tree
go build -o forge .
```

## Contributing

We welcome contributions of all kinds. To get started:

1. Fork the repository and create a new branch.
2. Make your changes, whether that's a bug fix, an improvement, or a new feature.
3. Open a pull request describing what you changed and why.
