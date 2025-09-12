package jira

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

// Client represents a Jira HTTP client
type Client struct {
	BaseURL    string
	APIToken   string
	HTTPClient *http.Client
}

// NewClient creates a new Jira client
func NewClient(baseURL, apiToken string) *Client {
	return &Client{
		BaseURL:  baseURL,
		APIToken: apiToken,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ValidateEnvironment checks that required environment variables are set
func ValidateEnvironment() error {
	jiraURL := os.Getenv("JIRA_URL")
	if jiraURL == "" {
		return fmt.Errorf("JIRA_URL environment variable is required")
	}
	
	apiToken := os.Getenv("JIRA_API_TOKEN")
	if apiToken == "" {
		return fmt.Errorf("JIRA_API_TOKEN environment variable is required")
	}
	
	// Validate URL format
	if _, err := url.Parse(jiraURL); err != nil {
		return fmt.Errorf("invalid JIRA_URL format: %w", err)
	}
	
	return nil
}

// ValidateIssueKey validates the issue key format
func ValidateIssueKey(issueKey string) error {
	// Pattern from specification: ^[A-Za-z]+-[0-9]+$
	pattern := `^[A-Za-z]+-[0-9]+$`
	matched, err := regexp.MatchString(pattern, issueKey)
	if err != nil {
		return fmt.Errorf("error validating issue key pattern: %w", err)
	}
	if !matched {
		return fmt.Errorf("invalid issue key format. Expected format: PROJECT-123")
	}
	return nil
}

// GetIssue fetches a single Jira issue by key
func (c *Client) GetIssue(ctx context.Context, issueKey string) (*Issue, error) {
	// Validate issue key format
	if err := ValidateIssueKey(issueKey); err != nil {
		return nil, err
	}
	
	// Construct URL
	endpoint := fmt.Sprintf("%s/rest/api/2/issue/%s", strings.TrimRight(c.BaseURL, "/"), issueKey)
	
	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	
	// Make request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		// Check for timeout
		if errors.Is(err, context.DeadlineExceeded) || strings.Contains(err.Error(), "timeout") {
			return nil, fmt.Errorf("network timeout after 30 seconds")
		}
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()
	
	// Handle different status codes per specification
	switch resp.StatusCode {
	case http.StatusOK:
		// Success - parse response
		var issue Issue
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}
		return &issue, nil
		
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed. Check JIRA_API_TOKEN environment variable")
		
	case http.StatusForbidden:
		return nil, fmt.Errorf("access denied to issue '%s'. Verify you have permission to view this issue", issueKey)
		
	case http.StatusNotFound:
		return nil, fmt.Errorf("issue '%s' not found. Verify the issue key exists and you have permission to view it", issueKey)
		
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("Jira server error. Try again later or contact your Jira administrator")
		
	default:
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
}