package jira

import (
	"fmt"
	"strings"
	"time"
)

// JiraTime is a custom time type that handles Jira's timestamp format
type JiraTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler for JiraTime
func (jt *JiraTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	str := strings.Trim(string(data), `"`)
	
	// Handle Jira's timestamp format: "2022-05-30T08:01:52.885+0000"
	// Convert "+0000" to "+00:00" for RFC3339 compatibility
	if strings.HasSuffix(str, "+0000") {
		str = str[:len(str)-5] + "+00:00"
	} else if strings.HasSuffix(str, "-0000") {
		str = str[:len(str)-5] + "-00:00"
	}
	
	// Try parsing with milliseconds first
	t, err := time.Parse("2006-01-02T15:04:05.000-07:00", str)
	if err != nil {
		// Try parsing without milliseconds as fallback
		t, err = time.Parse("2006-01-02T15:04:05-07:00", str)
		if err != nil {
			return fmt.Errorf("unable to parse time %q: %w", str, err)
		}
	}
	
	jt.Time = t
	return nil
}

// Issue represents a Jira issue response
type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

// IssueFields represents the fields section of a Jira issue
type IssueFields struct {
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Status      *Status     `json:"status"`
	IssueType   *IssueType  `json:"issuetype"`
	Priority    *Priority   `json:"priority"`
	Assignee    *User       `json:"assignee"`
	Reporter    *User       `json:"reporter"`
	Created     JiraTime    `json:"created"`
	Updated     JiraTime    `json:"updated"`
}

// Status represents a Jira issue status
type Status struct {
	Name string `json:"name"`
}

// IssueType represents a Jira issue type
type IssueType struct {
	Name string `json:"name"`
}

// Priority represents a Jira issue priority
type Priority struct {
	Name string `json:"name"`
}

// User represents a Jira user
type User struct {
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
}