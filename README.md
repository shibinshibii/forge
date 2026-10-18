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

## Generated structure

```text
my-api/
├── cmd/server/main.go
├── internal/config/config.go
├── internal/handlers/handler.go
├── migrations/
├── .gitignore
├── README.md
└── go.mod
```

## Develop

```sh
git clone https://github.com/shibinshibii/forge.git
cd forge
go run . create test-api
go build -o forge .
```

## Contributing

Bug reports and pull requests are welcome. Keep changes focused, run `gofmt` on Go files, and verify with `go test ./...` before opening a pull request.
