package data

type File struct {
	Name   string
	Folder string
	Data   []byte
}

var Dirs = []string{
	"cmd/server",
	"internal/config",
	"internal/handlers",
	"migrations",
}
