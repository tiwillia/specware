# Specware Implementation Plan

## Phase 1: Basic structure

Create the `specware` tool in golang. Use the cobra library. Use a directory structure that puts command-related files in cmd/ and implementation specifics in either internal/ or pkg/.

Add the sub-commands as specified in the spec, but they do nothing but output "Not yet implemented".

Add a simple README.md file.

Add a Makefile supporting `make build` that builds the tool.

## Phase 2: Template and Claude Command creation
Add the .claude/commands directory and command files with a basic summary/workflow structure.

Add the templates/ direcotry and create the necessary templates in this directory. They should remain simple for now, containing one-sentence descriptions of each section within the section and only a few sections.

## Phase 3: Init sub-command

Implement the `init` sub-command functionality to create the spec files and copy the commands into the specified directory.

Add testing using ginkgo to validate all functionality.

## Phase 4: Localized Templates Sub-command

Implement the `localize-templates` sub-command functionality to copy embedded templates to `.spec/templates/` directory for project-specific customization.

- Create `.spec/templates/` directory if it doesn't exist
- Copy embedded template files (requirements.md, implementation-plan.md) to the templates directory
- Handle cases where templates already exist (overwrite vs preserve)
- Add comprehensive testing using ginkgo to validate template copying functionality
- Test edge cases like missing .spec directory, permission issues, and existing template files

## Phase 5: Feature Sub-commands (new-requirements and new-implementation-plan)

Implement the core feature management sub-commands for creating specification artifacts.

### `feature new-requirements <short-name>`
- Generate sequential feature directory (e.g., `001-<short-name>/`)
- Copy requirements.md template (from localized templates if available, otherwise embedded)
- Create q&a-requirements.md file for tracking Q&A sessions
- Handle feature name validation and directory naming conventions
- Skip updating or creating .spec-status, this should be done in phase 6.

### `feature new-implementation-plan <short-name>`
- Validate that feature directory exists
- Copy implementation-plan.md template to existing feature directory
- Create q&a-implementation-plan.md file for implementation Q&A tracking
- Skip updating or creating .spec-status, this should be done in phase 6.

### Testing
- Add comprehensive ginkgo tests for both sub-commands
- Test sequential numbering logic
- Test template selection (localized vs embedded)
- Test error handling for missing directories, invalid names, etc.

## Phase 6: Feature Status Update Sub-command (TODO/BLOCKER)

Implement the `feature update-state <short-name> <status>` sub-command and the status tracking in feature new-requirements and feature new-implemenation-plan.

**BLOCKER**: Requires technical definition of:
- Valid status values and state transitions
- `.spec-status` file format and structure
- Status validation logic
- Integration with Claude Code workflow phases

**TODO**: Define status management requirements before implementation.
