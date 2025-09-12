package cmd_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"specware/cmd"
)

var _ = Describe("Jira Command Integration", func() {
	var (
		originalJiraURL   string
		originalAPIToken  string
	)

	BeforeEach(func() {
		// Save original environment variables
		originalJiraURL = os.Getenv("JIRA_URL")
		originalAPIToken = os.Getenv("JIRA_API_TOKEN")
	})

	AfterEach(func() {
		// Restore original environment variables
		if originalJiraURL != "" {
			os.Setenv("JIRA_URL", originalJiraURL)
		} else {
			os.Unsetenv("JIRA_URL")
		}
		
		if originalAPIToken != "" {
			os.Setenv("JIRA_API_TOKEN", originalAPIToken)
		} else {
			os.Unsetenv("JIRA_API_TOKEN")
		}
	})

	Describe("CreateJiraCommand", func() {
		It("should create a valid jira command", func() {
			jiraCmd := cmd.CreateJiraCommand()
			
			Expect(jiraCmd.Use).To(Equal("jira"))
			Expect(jiraCmd.Short).To(ContainSubstring("Jira integration"))
			Expect(jiraCmd.HasSubCommands()).To(BeTrue())
			
			subCommands := jiraCmd.Commands()
			Expect(len(subCommands)).To(Equal(1))
			Expect(subCommands[0].Use).To(ContainSubstring("get-issue"))
		})
	})

	Describe("Command structure validation", func() {
		It("should have correct command hierarchy", func() {
			jiraCmd := cmd.CreateJiraCommand()
			getIssueCmd, _, err := jiraCmd.Find([]string{"get-issue"})
			
			Expect(err).NotTo(HaveOccurred())
			Expect(getIssueCmd.Use).To(ContainSubstring("get-issue"))
			Expect(getIssueCmd.Args).NotTo(BeNil()) // Should have argument validation
		})
	})

	Describe("Help text validation", func() {
		It("should contain required environment variable documentation", func() {
			jiraCmd := cmd.CreateJiraCommand()
			getIssueCmd, _, err := jiraCmd.Find([]string{"get-issue"})
			
			Expect(err).NotTo(HaveOccurred())
			Expect(getIssueCmd.Long).To(ContainSubstring("JIRA_URL"))
			Expect(getIssueCmd.Long).To(ContainSubstring("JIRA_API_TOKEN"))
			Expect(getIssueCmd.Long).To(ContainSubstring("specware jira get-issue"))
		})
	})
})