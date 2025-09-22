# Context: Implementation Plan

## Questions & Answers
Questions and provided answers

### Q1: Should we add a new function to the spec package for Claude Code settings management or integrate it into InitProject?
**Answer:** Add new function (separation of concerns, easier testing)
Example: `UpdateClaudeSettings(targetDir string, autoYes bool) error`

### Q2: How should we structure the Claude Code settings JSON to handle existing permissions?
**Answer:** Certainly preserve existing user's settings. This should _never_ remove any existing settings.
```go
type ClaudeSettings struct {
    Permissions *PermissionsConfig `json:"permissions,omitempty"`
}
type PermissionsConfig struct {
    Allow []string `json:"allow,omitempty"`
}
```

### Q3: Should the prompt for permissions update happen before or after creating the project files?
**Answer:** After

### Q4: How should we handle user input for the permissions prompt?
**Answer:** Use bufio.Scanner with stdin (consistent with Go standards, handles whitespace)
```go
fmt.Print("Would you like to update Claude Code permissions? (y/N): ")
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
response := strings.ToLower(strings.TrimSpace(scanner.Text()))
```

### Q5: Should the flag be added to the init command specifically or support a global -y flag pattern?
**Answer:** Init command specifically. There is no use for it globally yet. YAGNI (You aren't going to need it)
```go
var yesFlag bool
initCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "automatically answer yes to all prompts")
```

### Q6: Should we write unit tests for the Claude Code settings functionality using Test-Driven Development?
**Answer:** Yes

### Q7: How should we test user interaction scenarios with the -y flag and manual prompts?
**Answer:** Use test helpers to mock stdin/stdout (isolate user input testing)
Example: Create helper functions to simulate user input and capture output

## Context Gathering Results
Collection of any additional detail needed to inform requirements.

### Current Codebase Structure
- **Command Structure**: Uses Cobra CLI framework with commands in `cmd/` directory
- **Init Command**: Located in `cmd/init.go`, calls `spec.InitProject()` from `internal/spec/spec.go`
- **Business Logic**: Core functionality in `internal/spec/spec.go` package
- **JSON Handling**: Already uses `encoding/json` for `.spec-status.json` files with `json.MarshalIndent()`
- **Error Handling**: Consistent pattern of returning `fmt.Errorf()` with wrapped errors

### Existing Patterns to Follow
1. **Command Definition**: Commands defined as `cobra.Command` structs with `Use`, `Short`, `Args`, and `Run` fields
2. **Flag Handling**: No existing flags in current commands, will need to add flag support
3. **File Operations**: Uses `os.WriteFile()`, `os.MkdirAll()`, and `filepath.Join()` consistently
4. **Error Messages**: Format: `"failed to <action>: %w"` with error wrapping
5. **Output Format**: Uses `fmt.Printf()` and `fmt.Println()` for user feedback
6. **Return Patterns**: Functions return `([]string, error)` for created files list and error

### JSON Structure Experience
- Codebase already handles JSON with structs using tags: `json:"current-step"`
- Uses `json.MarshalIndent(statusData, "", "  ")` for pretty formatting
- Validates JSON unmarshaling in tests

### File Path Patterns
- Consistent use of `filepath.Join()` for cross-platform compatibility
- Target directory passed as parameter to functions
- Relative paths stored in `createdFiles` slice for user feedback

### Testing Approach
- Uses Ginkgo v2 and Gomega for BDD-style testing in `internal/spec/spec_test.go`
- Tests create temporary directories and verify file creation
- JSON content validation through unmarshaling and struct comparison
