package assets

import (
	"embed"
	"io/fs"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed commands/*
var commandsFS embed.FS

//go:embed spec-readme.md
var specReadmeContent embed.FS

// GetTemplatesFS returns the embedded templates filesystem
func GetTemplatesFS() fs.FS {
	return templatesFS
}

// GetCommandsFS returns the embedded commands filesystem  
func GetCommandsFS() fs.FS {
	return commandsFS
}

// GetSpecReadme returns the content of the spec README file
func GetSpecReadme() ([]byte, error) {
	return specReadmeContent.ReadFile("spec-readme.md")
}