# Context: Implementation Plan

## Codebase Architecture Analysis

### Project Structure
- **Root structure**: Go module with Cobra CLI framework
- **Command layer**: `cmd/` package with root.go registering subcommands
- **Business logic**: `internal/spec/` package containing core functionality  
- **Assets**: `assets/` package with embedded filesystem for templates, commands, agents
- **Testing**: Ginkgo/Gomega BDD framework in `internal/spec/*_test.go`

### Command Implementation Patterns
- **Command definition**: Cobra commands in `cmd/` files with Use, Short, Long, Args, Run
- **Function delegation**: Commands call `internal/spec` functions with error handling
- **Error handling**: Commands print errors and call `os.Exit(1)` on failure
- **File operations**: Functions return `([]string, error)` where strings are created file paths
- **Status updates**: All commands follow pattern of status output then file listing

### Testing Patterns
- **Framework**: Ginkgo v2 with Gomega assertions in `spec_test` package
- **Setup**: `BeforeEach`/`AfterEach` with temp directory creation/cleanup
- **Tests**: Describe/It structure testing directory creation and file existence
- **Assertions**: `Expect(err).NotTo(HaveOccurred())` and `Expect(path).To(BeADirectory())`

### Existing Subcommand Structure
- **Root command**: `specware` in `cmd/root.go:22-25` adds: initCmd, localizeTemplatesCmd, featureCmd
- **Feature subcommands**: `cmd/feature.go:118-122` adds: newRequirementsCmd, newImplementationPlanCmd, updateStateCmd
- **Pattern**: Each command calls internal/spec function, prints results, handles errors

### Dependencies and Build
- **Core deps**: cobra CLI, ginkgo/gomega testing, embed FS
- **No HTTP libs**: Currently no HTTP client dependencies 
- **Build**: Simple `go build -o specware .` and `go test ./...` commands
- **Module**: Go 1.24.5 with standard semantic versioning

### Code Style Observations
- **Error messages**: Descriptive with context using fmt.Errorf
- **Validation**: Input validation functions (ValidateFeatureName) with regex patterns
- **File handling**: Standard os package operations with 0644/0755 permissions
- **JSON**: Standard encoding/json for .spec-status.json files
- **Package organization**: Clear separation between CLI layer and business logic

## Questions & Answers
Questions and provided answers

### Q1: Should the new Jira command be added to the root command or as a subcommand under 'feature'?
**Answer:** It should be added under a new sub-command `jira` like `specware jira get-issue`. This ensures scalability and seperates it from the feature commands and the core "project setup" commands at the root level.

### Q2: Should we create a new file `cmd/jira.go` following the existing pattern?
**Answer:** y

### Q3: Should the HTTP client implementation be in the internal/spec package or a new internal/jira package?
**Answer:** New internal jira package

### Q4: Should we use Go's standard net/http package or add an external HTTP client dependency?
**Answer:** You should review whether there is an official, well-documented, and often used golang SDK for jira. If that exists, we should prefer to use it. If it does not exist, standard packages are preferred.

### Q5: Should the output formatting logic be a separate function in internal/spec package?
**Answer:** It should be a seperate function in the internal/jira package

### Testing Questions

### Q6: Should we write unit tests for the HTTP client and formatting functions separately?
**Answer:** Yes

### Q7: Should we create integration tests that make actual HTTP calls to a test Jira instance?
**Answer:** No, we should create integration tests that make calls to a mock jira instance.

## Context Gathering Results
Collection of any additional detail needed to inform requirements.
