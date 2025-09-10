# specware - Spec-driven workflow tool

Spec-driven workflow enablement tool

## USAGE
```
specware <COMMAND> [OPTIONS]
```

## COMMANDS

### Core Commands
- `init <directory>` - Initialize project to support spec-driven-workflow
- `localize-templates` - Create project-specific templates

### Feature Specification
- `feature new-requirements <short-name>` - Create new feature specification directory
- `feature new-implementation-plan <short-name>` - Create implementation plan for existing feature
- `feature update-state <short-name> <status>` - Update the status of a feature specification

## EXAMPLES

```bash
# Initialize spec workflow in current directory
specware init .

# Create localized templates for customization
specware localize-templates

# Create new feature specification
specware feature new-requirements community-page

# Create implementation plan for existing feature
specware feature new-implementation-plan community-page

# Update feature status
specware feature update-state community-page "Requirements Q&A"
```

## WORKFLOW

1. Initialize project with `specware init .`
2. Use Claude Code command `/specify` to begin requirements gathering
3. Tool creates feature directories and templates automatically
4. Continue with implementation planning phase
5. Use generated artifacts to implement features

## DIRECTORY STRUCTURE

```
.claude/commands/specify.md    - Claude Code command for workflow
.spec/                         - Main specification directory
.spec/README.md               - Getting started documentation
.spec/NNN-feature-name/       - Individual feature directories
.spec/templates/              - Localized templates (optional)
```

For more information, see `.spec/README.md` after initialization.
