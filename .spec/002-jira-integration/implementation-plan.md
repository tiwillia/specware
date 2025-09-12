# Implementation Plan: Jira Integration for Specification Building

## Technical Approach

This implementation adds Jira integration capabilities to the specware CLI tool by creating a new `jira` subcommand with `get-issue` functionality. The implementation follows existing patterns:

- **New command structure**: `cmd/jira.go` following the pattern of `cmd/feature.go`
- **New business logic package**: `internal/jira/` for HTTP client and formatting logic
- **Cobra integration**: Add `jiraCmd` to root command registration
- **Environment-based configuration**: Use `JIRA_URL` and `JIRA_API_TOKEN` variables
- **Standard library preference**: Evaluate Go Jira SDK vs net/http package
- **Comprehensive testing**: Unit tests with mocked HTTP server for integration testing

The feature integrates seamlessly with existing architecture while maintaining separation of concerns through the new `internal/jira` package.

## Implementation

### Phase 1: Dependency Research and Project Setup
- [ ] Step 1: Research available Go Jira SDKs (andygrunwald/go-jira, trivago/tgo, others)
- [ ] Step 2: Evaluate SDK documentation, maintenance status, and feature fit
- [ ] Step 3: Make dependency decision (SDK vs standard library)
- [ ] Step 4: Update go.mod with selected dependency if needed
- [ ] Step 5: Run `go mod tidy` to clean dependencies
- [ ] Step 6: Run tests to ensure existing functionality unaffected: `make test`

### Phase 2: Create Internal Jira Package Structure
- [ ] Step 7: Create `internal/jira/` directory
- [ ] Step 8: Create `internal/jira/client.go` for HTTP client functionality
```go
package jira

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

type Client struct {
    BaseURL    string
    APIToken   string
    HTTPClient *http.Client
}

func NewClient(baseURL, apiToken string) *Client {
    return &Client{
        BaseURL:  baseURL,
        APIToken: apiToken,
        HTTPClient: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (c *Client) GetIssue(issueKey string) (*Issue, error) {
    // Implementation
}
```
- [ ] Step 9: Create `internal/jira/types.go` for Jira response structures
- [ ] Step 10: Create `internal/jira/formatter.go` for output formatting
- [ ] Step 11: Run tests to ensure package compiles: `go build ./internal/jira`

### Phase 3: Implement Core Jira Client Functionality
- [ ] Step 12: Implement environment variable validation in `client.go`
```go
func ValidateEnvironment() error {
    if os.Getenv("JIRA_URL") == "" {
        return fmt.Errorf("JIRA_URL environment variable is required")
    }
    if os.Getenv("JIRA_API_TOKEN") == "" {
        return fmt.Errorf("JIRA_API_TOKEN environment variable is required")
    }
    return nil
}
```
- [ ] Step 13: Implement HTTP request building with authentication
- [ ] Step 14: Implement JSON response parsing according to api-integration-spec.yaml
- [ ] Step 15: Implement error handling for all HTTP status codes per specification
- [ ] Step 16: Test HTTP client with mock server: `go test ./internal/jira`

### Phase 4: Implement Output Formatting
- [ ] Step 17: Implement `FormatIssue()` function following output-format-spec.txt
```go
func FormatIssue(issue *Issue) string {
    var output strings.Builder
    output.WriteString(fmt.Sprintf("Issue: %s\n", issue.Key))
    output.WriteString(fmt.Sprintf("Title: %s\n", getFieldOrDefault(issue.Fields.Summary, "No title provided")))
    // Continue per specification...
    return output.String()
}
```
- [ ] Step 18: Implement helper functions for null/empty field handling
- [ ] Step 19: Test formatting with various issue scenarios
- [ ] Step 20: Run formatting tests: `go test ./internal/jira -run TestFormatIssue`

### Phase 5: Create Jira Command Structure
- [ ] Step 21: Create `cmd/jira.go` following existing Cobra patterns
```go
package cmd

var jiraCmd = &cobra.Command{
    Use:   "jira",
    Short: "Jira integration commands",
}

var getIssueCmd = &cobra.Command{
    Use:   "get-issue <issue-key>",
    Short: "Fetch a single Jira issue",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation calling internal/jira
    },
}
```
- [ ] Step 22: Implement command validation and error handling
- [ ] Step 23: Add jiraCmd to root command in `cmd/root.go`
- [ ] Step 24: Test command registration: `go run . jira --help`

### Phase 6: Integrate Command with Jira Client
- [ ] Step 25: Wire command to call `internal/jira` functions
- [ ] Step 26: Implement environment variable validation at command startup
- [ ] Step 27: Implement issue key format validation per specification
- [ ] Step 28: Add proper error messages following existing patterns
- [ ] Step 29: Test end-to-end functionality: `go run . jira get-issue TEST-123`

### Phase 7: Comprehensive Testing Implementation
- [ ] Step 30: Create `internal/jira/client_test.go` with unit tests
- [ ] Step 31: Create mock HTTP server for integration testing
- [ ] Step 32: Test all error scenarios from api-integration-spec.yaml
- [ ] Step 33: Create `internal/jira/formatter_test.go` with output format tests
- [ ] Step 34: Test various field combinations and null values
- [ ] Step 35: Run full test suite: `make test`

### Phase 8: Final Integration and Validation
- [ ] Step 36: Build final binary: `make build`
- [ ] Step 37: Test with real Jira instance if available
- [ ] Step 38: Validate output format matches specification exactly
- [ ] Step 39: Test all error scenarios manually
- [ ] Step 40: Run final test suite to ensure no regressions: `make test`

## Summary of Changes in Key Technical Areas

### Components to Modify/Create

**New Components:**
- `cmd/jira.go` - Cobra command definition and CLI interface
- `internal/jira/client.go` - HTTP client and API communication
- `internal/jira/types.go` - Go structs for Jira API responses
- `internal/jira/formatter.go` - Issue output formatting logic
- `internal/jira/client_test.go` - Unit and integration tests for client
- `internal/jira/formatter_test.go` - Unit tests for formatting logic

**Modified Components:**
- `cmd/root.go` - Add jiraCmd to command registration
- `go.mod` - Add Jira SDK dependency if selected over standard library

### Database Changes
No database changes required - this is a CLI tool integration.

### API Changes
This feature creates a new CLI command but does not modify existing APIs. It consumes the Jira REST API `/rest/api/2/issue/{issueKey}` endpoint.

### User Interface Changes
Adds new CLI command structure:
```
specware jira get-issue <issue-key>
```

## Testing Strategy

### Unit Tests
- **Client functions**: Test HTTP request building, response parsing, error handling
- **Formatting functions**: Test output format with various input scenarios
- **Environment validation**: Test missing/empty environment variable handling
- **Issue key validation**: Test case-insensitive pattern matching

### Integration Tests
- **Mock HTTP server**: Test full HTTP request/response cycle with realistic Jira API responses
- **Error scenarios**: Test network timeouts, authentication failures, malformed responses
- **End-to-end command testing**: Test CLI command with mocked backend

### Test Execution
- Unit tests: `go test ./internal/jira`
- Full test suite: `make test`
- Manual testing: `make build && ./specware jira get-issue PROJ-123`

## Deployment Considerations

### Environment Configuration
Users must set environment variables:
```bash
export JIRA_URL="https://company.atlassian.net"
export JIRA_API_TOKEN="your-personal-access-token"
```

### Build and Distribution
- Standard Go build process: `make build`
- Binary distribution: Single `specware` executable
- No additional runtime dependencies beyond the binary

### Backward Compatibility
- No breaking changes to existing commands
- New functionality is additive only
- Existing users unaffected by this feature
