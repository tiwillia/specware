package spec_test

import (
	"os"
	"path/filepath"
	"specware/internal/spec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spec", func() {
	var tempDir string

	BeforeEach(func() {
		var err error
		tempDir, err = os.MkdirTemp("", "specware-test")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
	})

	Describe("InitProject", func() {
		It("should create the necessary directory structure", func() {
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())

			// Check .claude/commands directory exists
			claudeDir := filepath.Join(tempDir, ".claude", "commands")
			Expect(claudeDir).To(BeADirectory())

			// Check .spec directory exists
			specDir := filepath.Join(tempDir, ".spec")
			Expect(specDir).To(BeADirectory())

			// Check example spec directory exists
			exampleDir := filepath.Join(specDir, "000-example-spec")
			Expect(exampleDir).To(BeADirectory())
		})

		It("should copy the specify command file", func() {
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())

			specifyPath := filepath.Join(tempDir, ".claude", "commands", "specify.md")
			Expect(specifyPath).To(BeAnExistingFile())

			content, err := os.ReadFile(specifyPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("Specify - Spec-driven Development Workflow"))
		})

		It("should create .spec/README.md", func() {
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())

			readmePath := filepath.Join(tempDir, ".spec", "README.md")
			Expect(readmePath).To(BeAnExistingFile())

			content, err := os.ReadFile(readmePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("Spec-driven Development"))
		})

		It("should create .spec-status file in example directory", func() {
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())

			statusPath := filepath.Join(tempDir, ".spec", "000-example-spec", ".spec-status")
			Expect(statusPath).To(BeAnExistingFile())

			content, err := os.ReadFile(statusPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("status: initialized"))
		})
	})

	Describe("LocalizeTemplates", func() {
		It("should create templates directory", func() {
			err := spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			templatesDir := filepath.Join(tempDir, ".spec", "templates")
			Expect(templatesDir).To(BeADirectory())
		})

		It("should copy template files", func() {
			err := spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			requirementsPath := filepath.Join(tempDir, ".spec", "templates", "requirements.md")
			Expect(requirementsPath).To(BeAnExistingFile())

			planPath := filepath.Join(tempDir, ".spec", "templates", "implementation-plan.md")
			Expect(planPath).To(BeAnExistingFile())
		})
	})

	Describe("GetNextFeatureNumber", func() {
		It("should return 1 for non-existent directory", func() {
			num, err := spec.GetNextFeatureNumber("/nonexistent")
			Expect(err).NotTo(HaveOccurred())
			Expect(num).To(Equal(1))
		})

		It("should return 1 for empty directory", func() {
			specDir := filepath.Join(tempDir, ".spec")
			err := os.MkdirAll(specDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			num, err := spec.GetNextFeatureNumber(specDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(num).To(Equal(1))
		})

		It("should return next number for existing features", func() {
			specDir := filepath.Join(tempDir, ".spec")
			err := os.MkdirAll(specDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			// Create some feature directories
			err = os.MkdirAll(filepath.Join(specDir, "001-first-feature"), 0755)
			Expect(err).NotTo(HaveOccurred())
			err = os.MkdirAll(filepath.Join(specDir, "003-third-feature"), 0755)
			Expect(err).NotTo(HaveOccurred())

			num, err := spec.GetNextFeatureNumber(specDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(num).To(Equal(4))
		})
	})
})
