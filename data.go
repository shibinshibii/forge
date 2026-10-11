package main

var dirs = []string{"cmd/server", "internal/config", "internal/handlers", "migrations"}

var files = []File{

			{Name: "README.md",
				Folder: "",
				Data:   []byte(readmeContent),},
			{
				Name:   ".gitignore",
				Folder: "",
				Data:   []byte(gitignoreContent),
			},
			{
				Name:   "go.mod",
				Folder: "",
				Data:   []byte(gomodContent),
			},
		}