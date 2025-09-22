# Implementation Plan: Claude Code Settings Allowlist

## Technical Approach
This feature extends the existing `specware init` command to optionally update Claude Code project-level settings with an allowlist for specware commands. The implementation follows existing patterns in the codebase:

1. **Separation of Concerns**: Create a new function `UpdateClaudeSettings()` in the `spec` package to handle Claude Code settings management
2. **Command Enhancement**: Add `-y`/`--yes` flag to the init command using Cobra's flag system
3. **JSON Handling**: Use existing JSON patterns with `encoding/json` and `json.MarshalIndent()` for settings file manipulation
4. **Error Handling**: Follow existing error wrapping patterns with `fmt.Errorf()`
5. **User Experience**: Integrate into existing init workflow after project files are created

The feature preserves all existing settings and only adds the specware allowlist entry if the user opts in.

## Implementation

### Milestone 1: Core Claude Code Settings Functionality
Implement the Claude Code settings management functionality.

#### Phase 1: Claude Code Settings Implementation
- [ ] Step 1: Create Claude Code settings data structures in `internal/spec/spec.go`
```go
// ClaudeSettings represents the structure of .claude/settings.local.json
type ClaudeSettings struct {
    Permissions *PermissionsConfig `json:"permissions,omitempty"`
}

type PermissionsConfig struct {
    Allow []string `json:"allow,omitempty"`
}
```
- [ ] Step 2: Implement `UpdateClaudeSettings(targetDir string, autoYes bool) error` function
- [ ] Step 3: Add required imports: `bufio`, `os`, `strings` for user input handling
- [ ] Step 4: Add basic tests to existing `internal/spec/spec_test.go` for core functionality
- [ ] Step 5: Test implementation: `go test ./internal/spec`
- [ ] Step 6: Commit changes: `git add -A && git commit -m "Add Claude Code settings management functionality"`

### Milestone 2: Command Line Interface Integration
Integrate the settings functionality into the init command with flag support.

#### Phase 2: Command Flag Implementation
- [ ] Step 7: Add `-y`/`--yes` flag to `cmd/init.go`
```go
var yesFlag bool

func init() {
    initCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "automatically answer yes to all prompts")
}
```
- [ ] Step 8: Update init command's Run function to call settings update after project creation
- [ ] Step 9: Update help text in `cmd/init.go` Short/Long descriptions
```go
Short: "Initialize project to support spec-driven-workflow",
Long: `Initialize project to support spec-driven-workflow

This command creates the following directory structure:
  .claude/commands/     - Claude Code command files (includes /specify workflow)
  .claude/agents/       - Claude Code agent files for specialized workflows
  .spec/                - Feature specifications directory
  .spec/config.json     - Configuration for workflow question counts
  .spec/README.md       - Documentation for the spec workflow

Optional modifications (user will be prompted):
  .claude/settings.local.json - Updates project permissions to allow specware
                                commands without prompting (personal settings only)`,
```
- [ ] Step 10: Test manually: `./specware init test-dir --help`
- [ ] Step 11: Commit changes: `git add cmd/init.go && git commit -m "Add -y flag and help documentation to init command"`

#### Phase 3: Integration Testing
- [ ] Step 12: Create integration test suite in `tests/init_integration_test.go` using Ginkgo/Gomega
- [ ] Step 13: Test scenarios:
  - `specware init <dir>` with user saying "y" to prompt
  - `specware init <dir>` with user saying "n" to prompt
  - `specware init <dir> -y` (auto-yes)
  - Missing settings file scenario
  - Existing settings preservation
- [ ] Step 14: Verify `make test` includes new `tests/` directory (current `go test ./...` should work)
- [ ] Step 15: Run full test suite: `make test`
- [ ] Step 16: Test manually with built binary:
```bash
make build
./specware init test-project
./specware init test-project-auto -y
```
- [ ] Step 17: Commit test changes: `git add tests/ && git commit -m "Add integration tests for Claude Code settings feature"`

### Milestone 3: Final Validation and Documentation
Complete testing, build verification, and documentation updates.

#### Phase 4: End-to-End Validation
- [ ] Step 18: Run complete build and test cycle: `make clean && make build && make test`
- [ ] Step 19: Test with real Claude Code settings by creating a `.claude/settings.local.json` file:
```json
{
  "permissions": {
    "allow": ["Bash(git:*)"]
  }
}
```
- [ ] Step 20: Verify that running `./specware init test-merge -y` preserves existing settings and adds specware entry
- [ ] Step 21: Verify help output matches CLI specification: `./specware init --help`
- [ ] Step 22: Test idempotent behavior - run init twice on same directory
- [ ] Step 23: Final commit: `git add -A && git commit -m "Complete Claude Code settings allowlist feature implementation"`

## Summary of Changes in Key Technical Areas

### Components to Modify/Create
**New Components:**
- `UpdateClaudeSettings()` function in `internal/spec/spec.go`
- `ClaudeSettings` and `PermissionsConfig` structs
- Basic tests added to existing `internal/spec/spec_test.go`
- Integration test suite in `tests/init_integration_test.go` (new directory and file)

**Modified Components:**
- `cmd/init.go`: Add `-y` flag and help documentation
- `internal/spec/spec.go`: Import additional packages (`bufio`, `strings`)

### Database Changes
None. This feature only interacts with local JSON files.

### API Changes
None. This is a CLI-only feature with no external API dependencies.

### User Interface Changes
**Command Line Interface:**
- New `-y`/`--yes` flag for `specware init`
- Enhanced help documentation showing optional modifications
- New user prompt for Claude Code permissions update
- Improved feedback messages about settings modifications

## Testing Strategy

### Unit Testing
- **Basic Testing**: Add tests to existing test suite for core functionality
- **Coverage Areas**: JSON parsing, settings merging, error handling, file validation
- **Framework**: Ginkgo v2 and Gomega (consistent with existing tests)
- **Test File**: Add to existing `internal/spec/spec_test.go`

### Integration Testing
- **Architecture**: Use `tests/` directory for cross-package integration tests that span cmd and internal layers
- **Framework**: Ginkgo/Gomega for consistency with existing test patterns
- **Command Integration**: Test full `specware init` workflow with and without `-y` flag
- **User Interaction**: Test prompt responses and automatic flag behavior with stdin mocking
- **File System**: Test with existing and missing settings files using temporary directories
- **Settings Preservation**: Verify existing allowlist entries are maintained

### Manual Testing Scenarios
1. `./specware init new-project` - Interactive prompt testing
2. `./specware init auto-project -y` - Automatic flag testing
3. Test with existing `.claude/settings.local.json` containing other permissions
4. Test with malformed JSON settings file
5. Test help output: `./specware init --help`
6. Test idempotent behavior with repeated init commands

### Expected Test Output
- **Unit Tests**: Coverage of all error paths and success scenarios
- **Integration Tests**: End-to-end CLI behavior validation
- **Manual Tests**: User experience verification and edge case handling

## Deployment Considerations

### Build Requirements
- No additional dependencies beyond existing Go modules
- Compatible with existing `make build` command
- Existing `make test` (using `go test ./...`) will automatically include new `tests/` directory
- No changes to deployment infrastructure needed

### User Impact
- **Backward Compatibility**: Existing `specware init` behavior unchanged for users who decline permissions
- **Optional Feature**: All Claude Code settings modifications are opt-in only
- **Safe Defaults**: Default behavior preserves current user experience

### Rollback Plan
- Feature can be disabled by removing the prompt and flag handling
- No persistent state changes beyond optional local settings file modifications
- Users can manually edit `.claude/settings.local.json` to remove allowlist entries
