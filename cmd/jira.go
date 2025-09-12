package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"specware/internal/jira"
)

var jiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "Jira integration commands",
	Long:  `Commands for integrating with Jira to fetch issue information.`,
}

var getIssueCmd = &cobra.Command{
	Use:   "get-issue <issue-key>",
	Short: "Fetch a single Jira issue",
	Long: `Fetch and display a single Jira issue by its key.

Requires environment variables:
- JIRA_URL: The base URL of your Jira instance (e.g., https://company.atlassian.net)
- JIRA_API_TOKEN: Your personal access token with issue read permissions

Example:
  specware jira get-issue PROJ-123`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		issueKey := args[0]
		
		// Validate environment variables
		if err := jira.ValidateEnvironment(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		
		// Create client
		jiraURL := os.Getenv("JIRA_URL")
		apiToken := os.Getenv("JIRA_API_TOKEN")
		client := jira.NewClient(jiraURL, apiToken)
		
		// Fetch issue
		ctx := context.Background()
		issue, err := client.GetIssue(ctx, issueKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		
		// Format and display output
		output := jira.FormatIssue(issue)
		fmt.Print(output)
	},
}

func init() {
	jiraCmd.AddCommand(getIssueCmd)
}

// CreateJiraCommand creates a new jira command for testing
func CreateJiraCommand() *cobra.Command {
	// Return a copy of the main jira command for testing
	cmd := *jiraCmd
	cmd.ResetCommands()
	
	// Create a copy of the get-issue command
	getIssue := *getIssueCmd
	cmd.AddCommand(&getIssue)
	
	return &cmd
}