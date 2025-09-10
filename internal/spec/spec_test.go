package spec_test

import (
	"os"
	"path/filepath"
	"strings"
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

		It("should handle existing template files", func() {
			// First localization
			err := spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			requirementsPath := filepath.Join(tempDir, ".spec", "templates", "requirements.md")
			
			// Modify the existing file
			customContent := "# Custom Requirements Template\nThis has been modified"
			err = os.WriteFile(requirementsPath, []byte(customContent), 0644)
			Expect(err).NotTo(HaveOccurred())

			// Second localization should overwrite
			err = spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			// File should exist and have been overwritten
			Expect(requirementsPath).To(BeAnExistingFile())
			content, err := os.ReadFile(requirementsPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).NotTo(ContainSubstring("Custom Requirements Template"))
		})

		It("should work when .spec directory doesn't exist", func() {
			newTempDir, err := os.MkdirTemp("", "specware-no-spec")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(newTempDir)

			err = spec.LocalizeTemplates(newTempDir)
			Expect(err).NotTo(HaveOccurred())

			templatesDir := filepath.Join(newTempDir, ".spec", "templates")
			Expect(templatesDir).To(BeADirectory())

			requirementsPath := filepath.Join(templatesDir, "requirements.md")
			Expect(requirementsPath).To(BeAnExistingFile())
		})

		It("should handle permission errors gracefully", func() {
			if os.Getuid() == 0 {
				Skip("Skipping permission test when running as root")
			}

			// Create a read-only parent directory
			roDir := filepath.Join(tempDir, "readonly")
			err := os.MkdirAll(roDir, 0755)
			Expect(err).NotTo(HaveOccurred())
			
			err = os.Chmod(roDir, 0555) // read and execute only
			Expect(err).NotTo(HaveOccurred())
			defer os.Chmod(roDir, 0755) // restore permissions for cleanup

			err = spec.LocalizeTemplates(roDir)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to create templates directory"))
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

	Describe("ValidateFeatureName", func() {
		It("should accept valid feature names", func() {
			validNames := []string{"test-feature", "feature_123", "my-cool-feature", "abc123"}
			for _, name := range validNames {
				err := spec.ValidateFeatureName(name)
				Expect(err).NotTo(HaveOccurred(), "Name '%s' should be valid", name)
			}
		})

		It("should reject empty names", func() {
			err := spec.ValidateFeatureName("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cannot be empty"))
		})

		It("should reject names with invalid characters", func() {
			invalidNames := []string{"test feature", "test@feature", "test.feature", "test/feature"}
			for _, name := range invalidNames {
				err := spec.ValidateFeatureName(name)
				Expect(err).To(HaveOccurred(), "Name '%s' should be invalid", name)
				Expect(err.Error()).To(ContainSubstring("can only contain letters, numbers, hyphens, and underscores"))
			}
		})

		It("should reject names that are too long", func() {
			longName := strings.Repeat("a", 51)
			err := spec.ValidateFeatureName(longName)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("must be 50 characters or less"))
		})
	})

	Describe("CreateNewRequirements", func() {
		BeforeEach(func() {
			// Initialize project structure
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a new feature requirements directory", func() {
			err := spec.CreateNewRequirements(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())

			featureDir := filepath.Join(tempDir, ".spec", "001-test-feature")
			Expect(featureDir).To(BeADirectory())

			requirementsPath := filepath.Join(featureDir, "requirements.md")
			Expect(requirementsPath).To(BeAnExistingFile())

			qaPath := filepath.Join(featureDir, "q&a-requirements.md")
			Expect(qaPath).To(BeAnExistingFile())
		})

		It("should use sequential numbering for multiple features", func() {
			err := spec.CreateNewRequirements(tempDir, "first-feature")
			Expect(err).NotTo(HaveOccurred())

			err = spec.CreateNewRequirements(tempDir, "second-feature")
			Expect(err).NotTo(HaveOccurred())

			Expect(filepath.Join(tempDir, ".spec", "001-first-feature")).To(BeADirectory())
			Expect(filepath.Join(tempDir, ".spec", "002-second-feature")).To(BeADirectory())
		})

		It("should prefer localized templates when available", func() {
			// First localize templates
			err := spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			// Modify the localized template
			customContent := "# Custom Requirements Template\nThis is a customized template"
			localTemplatePath := filepath.Join(tempDir, ".spec", "templates", "requirements.md")
			err = os.WriteFile(localTemplatePath, []byte(customContent), 0644)
			Expect(err).NotTo(HaveOccurred())

			// Create new requirements
			err = spec.CreateNewRequirements(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())

			// Check that the customized template was used
			requirementsPath := filepath.Join(tempDir, ".spec", "001-test-feature", "requirements.md")
			content, err := os.ReadFile(requirementsPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("Custom Requirements Template"))
		})

		It("should fail with invalid feature names", func() {
			err := spec.CreateNewRequirements(tempDir, "")
			Expect(err).To(HaveOccurred())

			err = spec.CreateNewRequirements(tempDir, "invalid name")
			Expect(err).To(HaveOccurred())
		})

		It("should fail if .spec directory doesn't exist", func() {
			newTempDir, err := os.MkdirTemp("", "specware-no-spec")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(newTempDir)

			err = spec.CreateNewRequirements(newTempDir, "test-feature")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(".spec directory not found"))
		})
	})

	Describe("CreateNewImplementationPlan", func() {
		BeforeEach(func() {
			// Initialize project and create a feature
			err := spec.InitProject(tempDir)
			Expect(err).NotTo(HaveOccurred())
			err = spec.CreateNewRequirements(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create implementation plan for existing feature", func() {
			err := spec.CreateNewImplementationPlan(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())

			featureDir := filepath.Join(tempDir, ".spec", "001-test-feature")
			planPath := filepath.Join(featureDir, "implementation-plan.md")
			Expect(planPath).To(BeAnExistingFile())

			qaPath := filepath.Join(featureDir, "q&a-implementation-plan.md")
			Expect(qaPath).To(BeAnExistingFile())
		})

		It("should fail if feature doesn't exist", func() {
			err := spec.CreateNewImplementationPlan(tempDir, "nonexistent-feature")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("feature directory not found"))
		})

		It("should fail if implementation plan already exists", func() {
			err := spec.CreateNewImplementationPlan(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())

			// Try to create again
			err = spec.CreateNewImplementationPlan(tempDir, "test-feature")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("implementation plan already exists"))
		})

		It("should use localized templates when available", func() {
			// First localize templates
			err := spec.LocalizeTemplates(tempDir)
			Expect(err).NotTo(HaveOccurred())

			// Modify the localized template
			customContent := "# Custom Implementation Plan\nThis is a customized plan template"
			localTemplatePath := filepath.Join(tempDir, ".spec", "templates", "implementation-plan.md")
			err = os.WriteFile(localTemplatePath, []byte(customContent), 0644)
			Expect(err).NotTo(HaveOccurred())

			// Create implementation plan
			err = spec.CreateNewImplementationPlan(tempDir, "test-feature")
			Expect(err).NotTo(HaveOccurred())

			// Check that the customized template was used
			planPath := filepath.Join(tempDir, ".spec", "001-test-feature", "implementation-plan.md")
			content, err := os.ReadFile(planPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("Custom Implementation Plan"))
		})

		It("should fail with invalid feature names", func() {
			err := spec.CreateNewImplementationPlan(tempDir, "")
			Expect(err).To(HaveOccurred())

			err = spec.CreateNewImplementationPlan(tempDir, "invalid name")
			Expect(err).To(HaveOccurred())
		})
	})
})
