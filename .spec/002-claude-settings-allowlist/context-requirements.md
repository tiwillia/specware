# Context: Requirements

## Questions & Answers
Questions and provided answers

### Q1: Should this allowlist update be opt-in during `specware init`?
**Answer:** Yes, by default this change should not be made. The user should be prompted during specware init on whether the project permissions should be updated. There should also be a flag added `-y` or `--yes` that will answer yes to all prompt questions, allowing the user to opt-in to everything without being prompted.

### Q2: Should the feature work with both project-level and global Claude Code settings?
**Answer:** Only project level

### Q3: Should users be able to modify the allowlist after initial setup?
**Answer:** Sort of. If the user runs `init` again and they don't already have the permissions updated, it should still prompt and update permissions.

### Q4: Should the feature provide feedback about which commands are being added to the allowlist?
**Answer:** Yes, transparency is absolutely required.

### Q5: Should the feature handle cases where Claude Code settings file doesn't exist yet?
**Answer:** No. Claude Code updates very often and usually fairly quietly. If we initiate what we think is the right simple project settings now, its possible we break future implementations of claude code. If the .claude/settings.local.json doesn't exist yet - we should explain that to the user as part of the init output and move on.

## Context Gathering Results
Collection of any additional detail needed to inform requirements.

### Current Init Command Implementation
- Located in `cmd/init.go` and `internal/spec/spec.go`
- The `InitProject()` function creates `.claude/commands`, `.claude/agents`, and `.spec` directories
- Currently creates files but does not modify any Claude Code settings
- Returns a list of created files for user feedback
- Shows next steps to the user after completion

### Claude Code Settings Structure
Based on official documentation at https://docs.claude.com/en/docs/claude-code/settings:

**Settings File Locations:**
- Project settings: `.claude/settings.local.json` (personal, not checked into source control)
- Shared project settings: `.claude/settings.json` (team-wide, can be checked in)

**Permission Structure:**
```json
{
  "permissions": {
    "allow": [
      "Bash(specware feature update-state:*)"
    ]
  }
}
```

**Key Implementation Details:**
- Bash rules use prefix matching, not regex
- The format would be `"Bash(specware feature update-state:*)"` to allow all specware update-state commands
- Settings files are hierarchical (user -> project shared -> project local)
- `.claude/settings.local.json` should not be checked into source control

### Current Cobra Command Structure
- The init command uses `cobra.ExactArgs(1)` for the target directory
- Command flags would need to be added for `--update-project-permissions`
- The current implementation already shows detailed feedback about created files

### Required Specware Commands to Allowlist
Based on the spec workflow, these commands need permission:
- `specware feature update-state <short-name> <status>`
- Potentially `specware feature new-requirements` and `specware feature new-implementation-plan`

### Q6: Should the allowlist include all specware commands or just the specific ones that cause permission prompts?
**Answer:** All specware commands

### Q7: Should the feature use `.claude/settings.local.json` or `.claude/settings.json` for the allowlist?
**Answer:** Personal settings only - users can manually modify project-level settings if/when they deem the tool usage something they'd like to check in. Assume they do not wish to check in specifications.

### Q8: Should the feature validate the existing settings file structure before modifying it?
**Answer:** Yes, if the structure changes in the future we need to be able to avoid mangling settings.

### Q9: Should there be a way to remove or undo the allowlist settings later?
**Answer:** No
