# specware

Spec-driven workflow enablement tool for Claude Code AI Coding Assistant.

## Installation

```bash
make build
```

## Usage

```bash
# Initialize spec workflow in current directory
./specware init .

# Create localized templates for customization
./specware localize-templates

# Create new feature specification
./specware feature new-requirements community-page

# Create implementation plan for existing feature
./specware feature new-implementation-plan community-page

# Update feature status
./specware feature update-state community-page "Requirements Q&A"
```

## Commands

- `init <directory>` - Initialize project to support spec-driven-workflow
- `localize-templates` - Create project-specific templates
- `feature new-requirements <short-name>` - Create new feature specification directory
- `feature new-implementation-plan <short-name>` - Create implementation plan for existing feature
- `feature update-state <short-name> <status>` - Update the status of a feature specification

For more information, see `.spec/README.md` after initialization.