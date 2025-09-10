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
- Skip updating or creating .spec-status.json, this should be done in phase 6.

### `feature new-implementation-plan <short-name>`
- Validate that feature directory exists
- Copy implementation-plan.md template to existing feature directory
- Create q&a-implementation-plan.md file for implementation Q&A tracking
- Skip updating or creating .spec-status.json, this should be done in phase 6.

### Testing
- Add comprehensive ginkgo tests for both sub-commands
- Test sequential numbering logic
- Test template selection (localized vs embedded)
- Test error handling for missing directories, invalid names, etc.

## Phase 6: Feature Status Update Sub-command

Implement the `feature update-state <short-name> <status>` sub-command and the status tracking in feature new-requirements and feature new-implementation-plan.

### Status File Format
- Use `.spec-status.json` as the filename for status tracking
- JSON structure: `{"current-step": "status-value"}`
- Simple, minimal, and extensible format

### Suggested Status Values
These values should be used by the specify.md Claude command:
- `requirements-gathering` - Initial requirements creation phase
- `requirements-qa` - Q&A session for requirements clarification
- `requirements-review` - User review of requirements documentation
- `implementation-planning` - Implementation plan creation phase
- `implementation-qa` - Q&A session for implementation planning
- `implementation-review` - User review of implementation plan
- `specification-complete` - Both requirements and implementation plan finalized

### Implementation Details
- Create `.spec-status.json` file when running `feature new-requirements`
- Initialize with `{"current-step": "requirements-gathering"}`
- Update status via `feature update-state <short-name> <status>` command
- Handle cases where feature directory or status file doesn't exist
- Add comprehensive ginkgo tests for status management functionality

## Phase 7: Specify.md Command Enhancement

Enhance the `internal/assets/commands/specify.md` Claude Code command to implement the complete workflow as described in the requirements document.

### Current State Analysis
The existing specify.md command provides a basic workflow structure but needs significant enhancement to:
- Use the correct status values defined in Phase 6
- Implement proper state tracking with the specware tool
- Follow the exact workflow phases described in requirements.md
- Provide clear guidance for resuming work at any point
- Handle all edge cases and continuation scenarios

### Required Enhancements

#### Status Value Updates
- Replace hardcoded status strings with the suggested status values from Phase 6
- Use consistent status progression: `requirements-gathering` → `requirements-qa` → `requirements-review` → `implementation-planning` → `implementation-qa` → `implementation-review` → `specification-complete`
- Update state tracking calls throughout the workflow

#### Workflow Structure Improvements
- Align command workflow exactly with requirements.md examples (Phase 1 and Phase 2)
- Add proper state checking and resumption logic
- Implement Q&A session tracking with appropriate status updates
- Add expert research phase with codebase analysis
- Include proper handoff between requirements and implementation phases

#### Command Guidance Enhancements
- Add clear instructions for checking current feature status
- Provide resumption commands for each workflow phase
- Include error handling guidance when specware tool is unavailable
- Add examples of proper specware tool usage throughout

#### Template Integration
- Ensure command works with both embedded and localized templates
- Add guidance for template customization workflow
- Include proper file structure expectations

### Implementation Tasks
- Update status values throughout the command to match Phase 6 definitions
- Enhance workflow phases to match requirements.md specification exactly
- Add state checking and resumption logic for interrupted workflows
- Improve Q&A session management with proper status tracking
- Add comprehensive error handling and edge case management
- Include examples and clarifications for each workflow phase

### Testing and Validation
- Manual testing of complete workflow from start to finish
- Test resumption scenarios at each phase
- Validate integration with all specware tool commands
- Ensure command works with both embedded and localized templates

### Sub-agent Review
Use the scope-creep-craig agent to review the enhanced specify.md command to ensure:
- The workflow is concise and easily understood by an LLM
- No unnecessary complexity has been added beyond requirements
- All instructions are clear and actionable
- The command maintains focus on the core spec-driven workflow
- No feature creep has occurred during enhancement
