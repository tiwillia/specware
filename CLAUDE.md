# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Specware is a CLI tool that enables spec-driven development workflows for Claude Code AI Coding Assistant. It creates structured directories and templates to facilitate feature specification and implementation planning.

## Build and Development Commands

```bash
# Build the project
make build

# Run tests (uses Ginkgo/Gomega framework)
go test ./...

# Run tests for specific package
go test ./internal/spec

# Clean build artifacts
make clean

# Run the built binary
./specware <command>
```

## Architecture

The project follows a standard Go CLI structure using Cobra for command handling:

- **`main.go`** - Entry point that calls `cmd.Execute()`
- **`cmd/`** - Cobra command definitions
  - `root.go` - Root command setup and subcommand registration
  - `init.go` - Project initialization command
  - `localize_templates.go` - Template localization command
  - `feature.go` - Feature management commands
- **`internal/spec/`** - Core business logic for spec-driven workflows
  - Handles project initialization with `.claude/commands` and `.spec` directories
  - Manages embedded asset copying (commands and templates)
  - Provides feature numbering and directory structure management
- **`internal/assets/`** - Embedded file system containing templates and commands

## Key Functionality

The tool creates a standardized project structure:
- `.claude/commands/` - Claude Code command files
- `.spec/` - Feature specifications directory with numbered folders (e.g., `001-feature-name`)
- `.spec/templates/` - Localized templates for customization
- `.spec-status` files - Track feature development phases

## Testing

Uses Ginkgo v2 and Gomega for BDD-style testing. Test files are located in `internal/spec/` with the pattern `*_test.go` and `*_suite_test.go`.