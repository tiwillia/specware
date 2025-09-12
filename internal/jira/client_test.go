package jira_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"specware/internal/jira"
)

var _ = Describe("Client", func() {
	var (
		client *jira.Client
		server *httptest.Server
	)

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Default handler - will be overridden in tests
			w.WriteHeader(http.StatusOK)
		}))
		client = jira.NewClient(server.URL, "test-token")
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("ValidateEnvironment", func() {
		var originalJiraURL, originalAPIToken string

		BeforeEach(func() {
			originalJiraURL = os.Getenv("JIRA_URL")
			originalAPIToken = os.Getenv("JIRA_API_TOKEN")
		})

		AfterEach(func() {
			os.Setenv("JIRA_URL", originalJiraURL)
			os.Setenv("JIRA_API_TOKEN", originalAPIToken)
		})

		Context("when both environment variables are set", func() {
			BeforeEach(func() {
				os.Setenv("JIRA_URL", "https://test.atlassian.net")
				os.Setenv("JIRA_API_TOKEN", "test-token")
			})

			It("should return no error", func() {
				err := jira.ValidateEnvironment()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when JIRA_URL is missing", func() {
			BeforeEach(func() {
				os.Unsetenv("JIRA_URL")
				os.Setenv("JIRA_API_TOKEN", "test-token")
			})

			It("should return an error", func() {
				err := jira.ValidateEnvironment()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("JIRA_URL environment variable is required"))
			})
		})

		Context("when JIRA_API_TOKEN is missing", func() {
			BeforeEach(func() {
				os.Setenv("JIRA_URL", "https://test.atlassian.net")
				os.Unsetenv("JIRA_API_TOKEN")
			})

			It("should return an error", func() {
				err := jira.ValidateEnvironment()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("JIRA_API_TOKEN environment variable is required"))
			})
		})

		Context("when JIRA_URL has invalid format", func() {
			BeforeEach(func() {
				os.Setenv("JIRA_URL", "://invalid-url")
				os.Setenv("JIRA_API_TOKEN", "test-token")
			})

			It("should return an error", func() {
				err := jira.ValidateEnvironment()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid JIRA_URL format"))
			})
		})

		Context("when JIRA_URL is malformed", func() {
			BeforeEach(func() {
				os.Setenv("JIRA_URL", "ht\ttp://invalid")
				os.Setenv("JIRA_API_TOKEN", "test-token")
			})

			It("should return an error", func() {
				err := jira.ValidateEnvironment()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid JIRA_URL format"))
			})
		})
	})

	Describe("ValidateIssueKey", func() {
		Context("with valid issue keys", func() {
			DescribeTable("should accept valid formats",
				func(issueKey string) {
					err := jira.ValidateIssueKey(issueKey)
					Expect(err).NotTo(HaveOccurred())
				},
				Entry("uppercase project", "PROJ-123"),
				Entry("lowercase project", "proj-123"),
				Entry("mixed case project", "Proj-123"),
				Entry("another mixed case", "pRoJ-456"),
				Entry("long project name", "PROJECTNAME-999"),
				Entry("single letter project", "A-1"),
				Entry("single letter lowercase", "a-1"),
			)
		})

		Context("with invalid issue keys", func() {
			DescribeTable("should reject invalid formats",
				func(issueKey string) {
					err := jira.ValidateIssueKey(issueKey)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid issue key format"))
				},
				Entry("missing project", "123"),
				Entry("missing number", "PROJ-"),
				Entry("no dash", "PROJ123"),
				Entry("multiple dashes", "PROJ-123-456"),
				Entry("empty string", ""),
				Entry("special characters in project", "PR@J-123"),
				Entry("letters in number", "PROJ-12a"),
			)
		})
	})

	Describe("GetIssue", func() {
		Context("when the request is successful", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Expect(r.Method).To(Equal("GET"))
					Expect(r.URL.Path).To(Equal("/rest/api/2/issue/PROJ-123"))
					Expect(r.Header.Get("Authorization")).To(Equal("Bearer test-token"))
					Expect(r.Header.Get("Accept")).To(Equal("application/json"))
					Expect(r.Header.Get("Content-Type")).To(Equal("application/json"))

					issue := jira.Issue{
						Key: "PROJ-123",
						Fields: jira.IssueFields{
							Summary:     "Test issue",
							Description: "Test description",
							Status:      &jira.Status{Name: "In Progress"},
							IssueType:   &jira.IssueType{Name: "Bug"},
							Priority:    &jira.Priority{Name: "High"},
							Assignee:    &jira.User{DisplayName: "John Doe"},
							Reporter:    &jira.User{DisplayName: "Jane Smith"},
							Created:     jira.JiraTime{Time: time.Now()},
							Updated:     jira.JiraTime{Time: time.Now()},
						},
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(issue)
				}))
				client = jira.NewClient(server.URL, "test-token")
			})

			It("should return the issue", func() {
				ctx := context.Background()
				issue, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).NotTo(HaveOccurred())
				Expect(issue).NotTo(BeNil())
				Expect(issue.Key).To(Equal("PROJ-123"))
				Expect(issue.Fields.Summary).To(Equal("Test issue"))
				Expect(issue.Fields.Status.Name).To(Equal("In Progress"))
			})
		})

		Context("when the issue key is invalid", func() {
			It("should return a validation error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "invalid-key")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid issue key format"))
			})
		})

		Context("when authentication fails", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnauthorized)
				}))
				client = jira.NewClient(server.URL, "invalid-token")
			})

			It("should return an authentication error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("authentication failed"))
				Expect(err.Error()).To(ContainSubstring("JIRA_API_TOKEN"))
			})
		})

		Context("when access is denied", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusForbidden)
				}))
				client = jira.NewClient(server.URL, "test-token")
			})

			It("should return a permission error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("access denied"))
				Expect(err.Error()).To(ContainSubstring("PROJ-123"))
			})
		})

		Context("when issue is not found", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}))
				client = jira.NewClient(server.URL, "test-token")
			})

			It("should return a not found error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-999")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("not found"))
				Expect(err.Error()).To(ContainSubstring("PROJ-999"))
			})
		})

		Context("when server returns internal error", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
				client = jira.NewClient(server.URL, "test-token")
			})

			It("should return a server error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("server error"))
			})
		})

		Context("when response is malformed JSON", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.Write([]byte("invalid json"))
				}))
				client = jira.NewClient(server.URL, "test-token")
			})

			It("should return a parsing error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("failed to parse response"))
			})
		})

		Context("when network timeout occurs", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Simulate a slow response that exceeds timeout
					time.Sleep(2 * time.Second)
					w.WriteHeader(http.StatusOK)
				}))
				client = jira.NewClient(server.URL, "test-token")
				// Override timeout for testing
				client.HTTPClient.Timeout = 1 * time.Second
			})

			It("should return a timeout error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("timeout"))
			})
		})

		Context("when server is unreachable", func() {
			BeforeEach(func() {
				// Use an invalid URL that will cause network error
				client = jira.NewClient("http://invalid-url-that-does-not-exist.local", "test-token")
			})

			It("should return a network error", func() {
				ctx := context.Background()
				_, err := client.GetIssue(ctx, "PROJ-123")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("network error"))
			})
		})
	})
})