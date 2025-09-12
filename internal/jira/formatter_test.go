package jira_test

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"specware/internal/jira"
)

var _ = Describe("Formatter", func() {
	Describe("FormatIssue", func() {
		Context("with a complete issue", func() {
			It("should format all fields correctly", func() {
				issue := &jira.Issue{
					Key: "PROJ-123",
					Fields: jira.IssueFields{
						Summary:     "Fix authentication bug",
						Description: "Users are experiencing authentication failures",
						Status:      &jira.Status{Name: "In Progress"},
						IssueType:   &jira.IssueType{Name: "Bug"},
						Priority:    &jira.Priority{Name: "High"},
						Assignee:    &jira.User{DisplayName: "John Smith"},
						Reporter:    &jira.User{DisplayName: "Jane Doe"},
						Created:     jira.JiraTime{Time: time.Now()},
						Updated:     jira.JiraTime{Time: time.Now()},
					},
				}

				output := jira.FormatIssue(issue)

				expectedOutput := `Issue: PROJ-123
Title: Fix authentication bug
Type: Bug
Status: In Progress
Priority: High
Assignee: John Smith

Description:
Users are experiencing authentication failures`

				Expect(output).To(Equal(expectedOutput))
			})
		})

		Context("with missing optional fields", func() {
			It("should use default values", func() {
				issue := &jira.Issue{
					Key: "PROJ-456",
					Fields: jira.IssueFields{
						Summary:     "Test issue",
						Description: "",
						Status:      nil,
						IssueType:   nil,
						Priority:    nil,
						Assignee:    nil,
					},
				}

				output := jira.FormatIssue(issue)

				expectedOutput := `Issue: PROJ-456
Title: Test issue
Type: Unknown
Status: Unknown
Priority: None
Assignee: Not assigned

Description:
No description provided`

				Expect(output).To(Equal(expectedOutput))
			})
		})

		Context("with empty summary", func() {
			It("should use default title", func() {
				issue := &jira.Issue{
					Key: "PROJ-789",
					Fields: jira.IssueFields{
						Summary:     "",
						Description: "Some description",
						Status:      &jira.Status{Name: "Done"},
						IssueType:   &jira.IssueType{Name: "Story"},
						Priority:    &jira.Priority{Name: "Medium"},
						Assignee:    &jira.User{DisplayName: "Alice Brown"},
					},
				}

				output := jira.FormatIssue(issue)

				Expect(output).To(ContainSubstring("Title: No title provided"))
			})
		})

		Context("with whitespace-only fields", func() {
			It("should treat them as empty", func() {
				issue := &jira.Issue{
					Key: "PROJ-999",
					Fields: jira.IssueFields{
						Summary:     "   ",
						Description: "\t\n",
						Status:      &jira.Status{Name: "  "},
						IssueType:   &jira.IssueType{Name: ""},
						Priority:    &jira.Priority{Name: "Low"},
						Assignee:    &jira.User{DisplayName: "   "},
					},
				}

				output := jira.FormatIssue(issue)

				Expect(output).To(ContainSubstring("Title: No title provided"))
				Expect(output).To(ContainSubstring("Type: Unknown"))
				Expect(output).To(ContainSubstring("Status: Unknown"))
				Expect(output).To(ContainSubstring("Assignee: Not assigned"))
				Expect(output).To(ContainSubstring("Description:\nNo description provided"))
			})
		})

		Context("with multiline description", func() {
			It("should preserve formatting", func() {
				description := `This is a bug report.

Steps to reproduce:
1. Open the app
2. Click login
3. Enter credentials

Expected: Success
Actual: Error`

				issue := &jira.Issue{
					Key: "PROJ-100",
					Fields: jira.IssueFields{
						Summary:     "Login bug",
						Description: description,
						Status:      &jira.Status{Name: "Open"},
						IssueType:   &jira.IssueType{Name: "Bug"},
						Priority:    &jira.Priority{Name: "Critical"},
						Assignee:    &jira.User{DisplayName: "Developer One"},
					},
				}

				output := jira.FormatIssue(issue)

				Expect(output).To(ContainSubstring("Description:\n" + description))
			})
		})

		Context("with special characters in fields", func() {
			It("should preserve them without modification", func() {
				issue := &jira.Issue{
					Key: "PROJ-200",
					Fields: jira.IssueFields{
						Summary:     "Handle & < > \" ' characters",
						Description: "Test with & < > \" ' special chars",
						Status:      &jira.Status{Name: "In Review"},
						IssueType:   &jira.IssueType{Name: "Task"},
						Priority:    &jira.Priority{Name: "Low"},
						Assignee:    &jira.User{DisplayName: "O'Brien"},
					},
				}

				output := jira.FormatIssue(issue)

				Expect(output).To(ContainSubstring("Handle & < > \" ' characters"))
				Expect(output).To(ContainSubstring("Test with & < > \" ' special chars"))
				Expect(output).To(ContainSubstring("O'Brien"))
			})
		})

		Context("output format validation", func() {
			It("should match the exact specification format", func() {
				issue := &jira.Issue{
					Key: "TEST-42",
					Fields: jira.IssueFields{
						Summary:     "Sample issue",
						Description: "Sample description",
						Status:      &jira.Status{Name: "To Do"},
						IssueType:   &jira.IssueType{Name: "Epic"},
						Priority:    &jira.Priority{Name: "Highest"},
						Assignee:    &jira.User{DisplayName: "Test User"},
					},
				}

				output := jira.FormatIssue(issue)

				lines := strings.Split(output, "\n")
				Expect(lines[0]).To(Equal("Issue: TEST-42"))
				Expect(lines[1]).To(Equal("Title: Sample issue"))
				Expect(lines[2]).To(Equal("Type: Epic"))
				Expect(lines[3]).To(Equal("Status: To Do"))
				Expect(lines[4]).To(Equal("Priority: Highest"))
				Expect(lines[5]).To(Equal("Assignee: Test User"))
				Expect(lines[6]).To(Equal(""))
				Expect(lines[7]).To(Equal("Description:"))
				Expect(lines[8]).To(Equal("Sample description"))
				
				// No trailing newlines
				Expect(output).NotTo(HaveSuffix("\n\n"))
			})
		})
	})
})