package jira

import (
	"fmt"
	"strings"
)

// FormatIssue formats a Jira issue according to the output specification
func FormatIssue(issue *Issue) string {
	var output strings.Builder
	
	// Header line: Issue: {key}
	output.WriteString(fmt.Sprintf("Issue: %s\n", issue.Key))
	
	// Title: {summary}
	summary := getFieldOrDefault(issue.Fields.Summary, "No title provided")
	output.WriteString(fmt.Sprintf("Title: %s\n", summary))
	
	// Type: {issuetype}
	issueType := "Unknown"
	if issue.Fields.IssueType != nil {
		issueType = getFieldOrDefault(issue.Fields.IssueType.Name, "Unknown")
	}
	output.WriteString(fmt.Sprintf("Type: %s\n", issueType))
	
	// Status: {status}
	status := "Unknown"
	if issue.Fields.Status != nil {
		status = getFieldOrDefault(issue.Fields.Status.Name, "Unknown")
	}
	output.WriteString(fmt.Sprintf("Status: %s\n", status))
	
	// Priority: {priority}
	priority := "None"
	if issue.Fields.Priority != nil {
		priority = getFieldOrDefault(issue.Fields.Priority.Name, "None")
	}
	output.WriteString(fmt.Sprintf("Priority: %s\n", priority))
	
	// Assignee: {assignee}
	assignee := "Not assigned"
	if issue.Fields.Assignee != nil {
		assignee = getFieldOrDefault(issue.Fields.Assignee.DisplayName, "Not assigned")
	}
	output.WriteString(fmt.Sprintf("Assignee: %s\n", assignee))
	
	// Blank line before description
	output.WriteString("\n")
	
	// Description section
	output.WriteString("Description:\n")
	description := getFieldOrDefault(issue.Fields.Description, "No description provided")
	output.WriteString(description)
	output.WriteString("\n")
	
	return output.String()
}

// getFieldOrDefault returns the field value if non-empty, otherwise returns the default
func getFieldOrDefault(field, defaultValue string) string {
	if strings.TrimSpace(field) == "" {
		return defaultValue
	}
	return field
}