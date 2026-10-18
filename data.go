
package main

type File struct {
	Name   string
	Folder string
	Data   []byte
}

var dirs = []string{
	"cmd/server",
	"internal/config",
	"internal/handlers",
	"migrations",
}
