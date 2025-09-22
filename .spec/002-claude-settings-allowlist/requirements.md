# Requirements Specification: Claude Code Settings Allowlist

## Problem Statement
Currently, when `specware` commands are executed to update specification status, Claude Code asks for permission every time a `specware feature update-state` command is run. This creates friction in the spec-driven development workflow and interrupts the development process.

## Solution Overview
Add an optional opt-in feature during `specware init` that updates the project-level Claude Code settings to add `specware` to the allow list. This will ensure Claude Code does not ask for permission every time an update to specification status is made, creating a smoother workflow experience.

## Functional Requirements

### Core Functionality
1. **Opt-in Prompt**: During `specware init`, prompt the user whether to update Claude Code project permissions
2. **Flag Support**: Support `-y`/`--yes` flag to automatically answer yes to all prompts, including the permissions update
3. **Settings Update**: When user opts in, update `.claude/settings.local.json` to allow all specware commands
4. **Idempotent Behavior**: If `specware init` is run again and permissions aren't set, offer to configure them again
5. **Transparency**: Clearly inform the user what permissions are being added and which file is being modified

### Permission Configuration
1. **Allowlist Scope**: Add allowlist entry for all specware commands: `"Bash(specware:*)"`
2. **Settings File**: Only modify `.claude/settings.local.json` (personal settings, not checked into source control)
3. **Structure Validation**: Validate existing settings file structure before modification to prevent corruption
4. **Graceful Handling**: If `.claude/settings.local.json` doesn't exist, explain this to the user and skip the update

### User Experience
1. **Clear Messaging**: Explain what the permission update does and why it's beneficial
2. **User Control**: Default behavior is NOT to update permissions (user must opt-in)
3. **Feedback**: Provide clear feedback about what was done or why it was skipped
4. **Help Documentation**: Update `specware init --help` to clearly document all modifications (required and optional)

## Technical Requirements

### Performance
- Settings file operations must be atomic to prevent corruption

### Security
- Only modify `.claude/settings.local.json`, never create it if it doesn't exist
- Validate JSON structure before modification
- Only allowlist trusted operations

### Compatibility
- Must work with existing Claude Code settings file formats
- Should be resilient to future Claude Code settings structure changes
- Must not break existing settings when updating

### Error Handling
- Handle missing settings file gracefully
- Handle malformed JSON gracefully
- Provide clear error messages for any failures

## Acceptance Criteria

1. **Basic Functionality**
   - [ ] `specware init <dir>` prompts user about updating permissions
   - [ ] `specware init <dir> -y` automatically updates permissions without prompting
   - [ ] Permissions are only updated when user explicitly opts in

2. **Settings Management**
   - [ ] Successfully adds `"Bash(specware:*)"` to `.claude/settings.local.json` allowlist
   - [ ] Preserves existing settings when updating
   - [ ] Handles missing settings file by informing user and continuing
   - [ ] Validates settings file structure before modification

3. **User Experience**
   - [ ] Clear explanation of what permissions will be added
   - [ ] Transparent feedback about what was modified
   - [ ] Graceful messaging when settings file doesn't exist
   - [ ] Idempotent behavior on repeated init runs

4. **Error Handling**
   - [ ] Continues gracefully if settings file is malformed
   - [ ] Provides helpful error messages for any failures
   - [ ] Does not corrupt existing settings on error

5. **CLI Interface**
   - [ ] `specware init --help` documents all files/directories created
   - [ ] Help clearly identifies optional modifications and prompting behavior
   - [ ] `-y`/`--yes` flag documented as "automatically answer yes to all prompts"
   - [ ] Help explains what `.claude/settings.local.json` modification does

## Constraints

### Technical Limitations
- Must not create `.claude/settings.local.json` if it doesn't exist (per user requirement)
- Cannot modify global Claude Code settings, only project-level personal settings
- Must be compatible with Cobra CLI framework used by specware

### Business Limitations
- Feature scope limited to specware commands only
- No undo functionality required (users can manually edit settings)
- Personal settings only (users can manually move to shared settings if desired)

### Dependencies
- Requires existing Claude Code installation and settings structure
- Depends on Claude Code maintaining current settings file format
- Requires Go JSON marshaling/unmarshaling for settings manipulation

