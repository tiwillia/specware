# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Specware is a CLI tool that enables spec-driven development workflows for Claude Code AI Coding Assistant. It creates structured directories and templates to facilitate feature specification and implementation planning with integrated status tracking.

## Build and Development Commands

```bash
# Build the project
make build

# Run all tests (unit and integration suites)
make test

# Run tests directly with go (unit and integration)
go test ./...

# Run specific test suites
go test ./internal/spec  # Unit tests only
go test ./tests          # Integration tests only

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
  - `feature.go` - Feature management commands with status tracking
- **`internal/spec/`** - Core business logic for spec-driven workflows
  - Handles project initialization with `.claude/commands`, `.claude/agents` and `.spec` directories
  - Manages embedded asset copying (commands, agents and templates)
  - Provides feature numbering and directory structure management
  - Implements status tracking through `.spec-status.json` files
- **`assets/`** - Embedded file system containing templates, commands and agents

## Key Functionality

The tool creates a standardized project structure:
- `.claude/commands/` - Claude Code command files (includes `specify.md` workflow)
- `.claude/agents/` - Claude Code agent files for specialized workflows
- `.spec/` - Feature specifications directory with numbered folders (e.g., `001-feature-name`)
- `.spec/templates/` - Localized templates for customization
- `.spec-status.json` files - Track feature development phases with JSON status

## Core Commands

- `specware init <directory>` - Initialize project with spec workflow support
- `specware localize-templates` - Copy embedded templates to `.spec/templates/` for customization
- `specware feature new-requirements <short-name>` - Create new feature specification directory
- `specware feature new-implementation-plan <short-name>` - Add implementation plan to existing feature
- `specware feature update-state <short-name> <status>` - Update feature status tracking

## Status Management

Features are tracked through `.spec-status.json` files with suggested statuses:
- `"Requirements Gathering"`
- `"Requirements Context Gathering"`
- `"Requirements Expert Q&A"`
- `"Requirements Complete"`
- `"Requirements Interactive Review"`
- `"Implementation Planning"`
- `"Implementation Plan Q&A"`
- `"Implementation Plan Generated"`
- `"Implementation Plan Interactive Review"`
- `"Implementation Planning Complete"`

## Testing

Uses Ginkgo v2 and Gomega for BDD-style testing with two test suites:

### Unit Tests
Located in `internal/spec/` with the pattern `*_test.go` and `*_suite_test.go`. These test individual components and business logic.

### Integration Tests
Located in `tests/` directory with full end-to-end testing of the CLI tool:
- `tests/integration_suite_test.go` - Integration test suite setup
- `tests/init_integration_test.go` - Tests for project initialization workflows including Claude Code settings integration
- Tests build the actual binary and execute commands in isolated temporary directories
- Validates complete workflows including allowlist updates and settings file handling