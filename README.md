# Specware

A CLI tool that enables spec-driven development workflows for the Claude Code AI Coding Assistant.

Specware creates structured directories, commands, and templates and provides a tool to Claude Code to facilitate feature requirements gathering and implementation planning.

## Quick Start

### 1. Initialize Your Project
```bash
# Initialize spec workflow in current directory
./specware init .

# Optional: Create localized templates for customization
./specware localize-templates
```

### 2. Start Feature Specification with Claude Code
```bash
# Open Claude Code and begin the spec-driven workflow
claude

# Use the /specify command to start gathering requirements
> /specify Add user authentication with email and password
```

The `/specify` command will guide you through:
- **Requirements gathering** - Interactive Q&A to understand the feature
- **Expert research** - Codebase analysis and technical planning  
- **Implementation planning** - Detailed technical specifications
- **Status tracking** - Progress monitoring throughout development

Claude Code automatically uses the `specware` tool to create directories, track progress, and maintain organized specifications for your features.

## Specware Tool Commands

### Project Setup
These commands are intended to be run by a user:
- `init <directory>` - Initialize project with spec-driven workflow support
- `localize-templates` - Copy embedded templates to `.spec/templates/` for customization, not required.

### Feature Management
These commands are intended to be run by Claude Code to facilitate feature specification:
- `feature new-requirements <short-name>` - Create new feature specification directory with requirements template
- `feature new-implementation-plan <short-name>` - Add implementation plan to existing feature
- `feature update-state <short-name> <status>` - Update feature development status

## Tooling Structure

After initialization, Specware creates in your project:

```
.claude/
  commands/
    specify.md          # Claude Code workflow command
.spec/
  README.md            # Spec workflow documentation  
  templates/           # Customizable templates (after localize-templates)
  001-feature-name/    # Feature directories with sequential numbering
    requirements.md
    implementation-plan.md
    q&a-requirements.md
    q&a-implementation-plan.md
    .spec-status.json   # Status tracking
```

## Status Tracking

Features are tracked through `.spec-status.json` files within a feature spec directory with suggested workflow phases:
- `requirements-gathering`
- `requirements-qa` 
- `requirements-review`
- `implementation-planning`
- `implementation-qa`
- `implementation-review`
- `specification-complete`
