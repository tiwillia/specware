package assets

import "embed"

//go:embed commands
var CommandsFS embed.FS

//go:embed agents
var AgentsFS embed.FS

//go:embed templates
var TemplatesFS embed.FS

//go:embed spec-readme.md
var SpecReadmeContent embed.FS

//go:embed config
var ConfigFS embed.FS