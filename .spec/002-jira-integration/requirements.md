# Requirements Specification: Jira Integration for Specification Building

## Problem Statement
Currently, the specware tool requires manual input for feature requirements gathering. When feature specifications are based on existing Jira issues, users must manually copy and format Jira issue content into the specification workflow. This creates friction and potential for errors when translating Jira issue details into specification requirements.

## Solution Overview
Add a Jira integration subcommand to the specware CLI tool that fetches issue data from Jira and outputs it in a format suitable for Claude Code consumption during specification building. This enables seamless integration of existing Jira issue context into the spec-driven development workflow.

## Functional Requirements

### Core Functionality
- Add `specware jira get-issue <issue-key>` subcommand to fetch a single Jira issue
- Read Jira URL and API token from environment variables `JIRA_URL` and `JIRA_API_TOKEN`
- Fetch issue data using Jira REST API `/rest/api/2/issue/{issueKey}` endpoint
- Output formatted text containing issue summary, description, status, and other standard fields
- Support authentication via personal access token only (no interactive authentication)
- Accept case-insensitive issue keys (minimal validation only)

### Environment Variable Handling
- Validate that `JIRA_URL` and `JIRA_API_TOKEN` environment variables exist and are non-empty strings
- Fail fast with clear error messages if environment variables are missing or empty
- Do not perform complex validation beyond non-empty string checks

### Output Format
- Format issue data as human-readable text (not JSON)
- Include standard Jira fields: key, summary, description, status, issue type, priority, assignee
- Preserve original issue content without modification
- Format output for optimal Claude Code readability and processing
- No emojis in output
- Use structured template format with field labels (see output-format-spec.txt)
- Handle missing/null fields with appropriate default values

### Error Handling
- Handle network connectivity issues with appropriate error messages
- Handle authentication failures with clear guidance
- Handle invalid issue keys with specific error messages
- Use 30-second timeout with no retry logic

## Technical Requirements

### Performance
- Single issue fetching only (no bulk operations)
- 30-second network timeout limit
- No local caching of issue data

### Security
- Use environment variables for sensitive authentication data
- Support personal access token authentication only
- Do not log or expose authentication tokens

### Compatibility
- Support single Jira instance per environment configuration
- Compatible with Jira Cloud and Server REST API v2
- Maintain existing specware CLI patterns and conventions

## Acceptance Criteria

### Command Structure
- [ ] `specware jira get-issue <issue-key>` command exists and is functional
- [ ] Command validates issue key parameter is provided
- [ ] Command follows existing specware error handling patterns

### Environment Variables
- [ ] `JIRA_URL` environment variable is required and validated as non-empty
- [ ] `JIRA_API_TOKEN` environment variable is required and validated as non-empty
- [ ] Clear error messages when environment variables are missing

### API Integration
- [ ] Successfully authenticates to Jira using personal access token
- [ ] Fetches issue data using Jira REST API `/rest/api/2/issue/{issueKey}` endpoint
- [ ] Handles API errors gracefully with user-friendly messages

### Output Format
- [ ] Outputs formatted text suitable for Claude Code processing
- [ ] Includes issue key, summary, description, status, type, priority, assignee
- [ ] Content is unmodified from Jira (preserves original formatting)
- [ ] Output is human-readable and well-structured
- [ ] No emojis in output

### Error Scenarios
- [ ] Handles missing environment variables
- [ ] Handles invalid Jira URLs
- [ ] Handles authentication failures
- [ ] Handles network connectivity issues
- [ ] Handles invalid or non-existent issue keys

## Constraints

### Scope Limitations
- Single issue fetching only (no multiple issues, linked issues, or sub-tasks)
- Standard Jira fields only (no custom field support)
- Personal access token authentication only
- Single Jira instance support per environment

### Technical Constraints
- Must follow existing specware CLI architecture patterns
- Must use Go standard library where possible
- Must maintain consistency with existing command structure
- No additional dependencies beyond necessary HTTP client functionality

### Dependencies
- Jira instance with REST API v2 support
- Valid personal access token with issue read permissions
- Network connectivity to Jira instance
- Go HTTP client library for API requests

### Technical Specifications
- API Integration: See api-integration-spec.yaml for detailed API contract
- Output Format: See output-format-spec.txt for exact formatting requirements

