package tests

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tiwillia/specware/internal/spec"
)

var _ = Describe("Init Integration Tests", func() {
	var tempDir string
	var specwareBinary string

	BeforeEach(func() {
		var err error
		tempDir, err = os.MkdirTemp("", "specware-integration-test")
		Expect(err).NotTo(HaveOccurred())

		// Build the specware binary for testing
		specwareBinary = filepath.Join(tempDir, "specware")
		// Get current working directory to build from the correct location
		wd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		projectRoot := filepath.Dir(wd) // Go up one level from tests/ to project root

		cmd := exec.Command("go", "build", "-o", specwareBinary, ".")
		cmd.Dir = projectRoot
		err = cmd.Run()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
	})

	Describe("specware init with Claude Code settings", func() {
		var testProjectDir string

		BeforeEach(func() {
			testProjectDir = filepath.Join(tempDir, "test-project")
		})

		It("should handle missing settings file gracefully with -y flag", func() {
			// Run specware init with auto-yes flag
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Successfully initialized spec-driven workflow"))
			Expect(outputStr).To(ContainSubstring("Claude Code settings file not found"))

			// Verify project structure was created
			Expect(filepath.Join(testProjectDir, ".claude", "commands")).To(BeADirectory())
			Expect(filepath.Join(testProjectDir, ".spec")).To(BeADirectory())
		})

		It("should update existing settings file with -y flag", func() {
			// Create .claude directory and settings file
			claudeDir := filepath.Join(testProjectDir, ".claude")
			err := os.MkdirAll(claudeDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			// Create existing settings with some permissions
			existingSettings := spec.ClaudeSettings{
				Permissions: &spec.PermissionsConfig{
					Allow: []string{"Bash(git:*)"},
				},
			}
			settingsData, err := json.MarshalIndent(existingSettings, "", "  ")
			Expect(err).NotTo(HaveOccurred())

			settingsPath := filepath.Join(claudeDir, "settings.local.json")
			err = os.WriteFile(settingsPath, settingsData, 0644)
			Expect(err).NotTo(HaveOccurred())

			// Run specware init with auto-yes flag
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Successfully initialized spec-driven workflow"))
			Expect(outputStr).To(ContainSubstring("Successfully updated Claude Code permissions"))
			Expect(outputStr).To(ContainSubstring("Added: \"" + spec.SpecwareAllowlistEntry + "\""))

			// Verify settings were updated correctly
			updatedData, err := os.ReadFile(settingsPath)
			Expect(err).NotTo(HaveOccurred())

			var updatedSettings spec.ClaudeSettings
			err = json.Unmarshal(updatedData, &updatedSettings)
			Expect(err).NotTo(HaveOccurred())

			Expect(updatedSettings.Permissions.Allow).To(ContainElement("Bash(git:*)"))
			Expect(updatedSettings.Permissions.Allow).To(ContainElement(spec.SpecwareAllowlistEntry))
		})

		It("should be idempotent - not add duplicate entries with -y flag", func() {
			// Create .claude directory and settings file with specware already present
			claudeDir := filepath.Join(testProjectDir, ".claude")
			err := os.MkdirAll(claudeDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			existingSettings := spec.ClaudeSettings{
				Permissions: &spec.PermissionsConfig{
					Allow: []string{spec.SpecwareAllowlistEntry, "Bash(git:*)"},
				},
			}
			settingsData, err := json.MarshalIndent(existingSettings, "", "  ")
			Expect(err).NotTo(HaveOccurred())

			settingsPath := filepath.Join(claudeDir, "settings.local.json")
			err = os.WriteFile(settingsPath, settingsData, 0644)
			Expect(err).NotTo(HaveOccurred())

			// Run specware init with auto-yes flag
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Successfully initialized spec-driven workflow"))
			Expect(outputStr).To(ContainSubstring("Specware permissions already configured"))

			// Verify no duplicates were added
			updatedData, err := os.ReadFile(settingsPath)
			Expect(err).NotTo(HaveOccurred())

			var updatedSettings spec.ClaudeSettings
			err = json.Unmarshal(updatedData, &updatedSettings)
			Expect(err).NotTo(HaveOccurred())

			// Count occurrences of specware entry
			count := 0
			for _, entry := range updatedSettings.Permissions.Allow {
				if entry == spec.SpecwareAllowlistEntry {
					count++
				}
			}
			Expect(count).To(Equal(1))
		})

		It("should handle malformed JSON gracefully with -y flag", func() {
			// Create .claude directory and malformed settings file
			claudeDir := filepath.Join(testProjectDir, ".claude")
			err := os.MkdirAll(claudeDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			settingsPath := filepath.Join(claudeDir, "settings.local.json")
			err = os.WriteFile(settingsPath, []byte("{ invalid json"), 0644)
			Expect(err).NotTo(HaveOccurred())

			// Run specware init with auto-yes flag
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Successfully initialized spec-driven workflow"))
			Expect(outputStr).To(ContainSubstring("Warning: Settings file appears to be malformed JSON"))

			// Verify malformed file was not modified
			content, err := os.ReadFile(settingsPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(Equal("{ invalid json"))
		})

		It("should work without -y flag on missing settings file", func() {
			// Run specware init without auto-yes flag (will skip settings update due to missing file)
			cmd := exec.Command(specwareBinary, "init", testProjectDir)
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Successfully initialized spec-driven workflow"))
			Expect(outputStr).To(ContainSubstring("Claude Code settings file not found"))

			// Verify project structure was still created
			Expect(filepath.Join(testProjectDir, ".claude", "commands")).To(BeADirectory())
			Expect(filepath.Join(testProjectDir, ".spec")).To(BeADirectory())
		})
	})

	Describe("specware init --help", func() {
		It("should show updated help documentation", func() {
			cmd := exec.Command(specwareBinary, "init", "--help")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Initialize project to support spec-driven-workflow"))
			Expect(outputStr).To(ContainSubstring(".claude/commands/"))
			Expect(outputStr).To(ContainSubstring(".claude/agents/"))
			Expect(outputStr).To(ContainSubstring(".spec/"))
			Expect(outputStr).To(ContainSubstring("Optional modifications (user will be prompted):"))
			Expect(outputStr).To(ContainSubstring(".claude/settings.local.json"))
			Expect(outputStr).To(ContainSubstring("-y, --yes"))
			Expect(outputStr).To(ContainSubstring("automatically answer yes to all prompts"))
		})
	})

	Describe("Repeated init runs", func() {
		var testProjectDir string

		BeforeEach(func() {
			testProjectDir = filepath.Join(tempDir, "test-project")

			// Initial setup - create project and settings
			claudeDir := filepath.Join(testProjectDir, ".claude")
			err := os.MkdirAll(claudeDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			initialSettings := spec.ClaudeSettings{
				Permissions: &spec.PermissionsConfig{
					Allow: []string{"Bash(git:*)"},
				},
			}
			settingsData, err := json.MarshalIndent(initialSettings, "", "  ")
			Expect(err).NotTo(HaveOccurred())

			settingsPath := filepath.Join(claudeDir, "settings.local.json")
			err = os.WriteFile(settingsPath, settingsData, 0644)
			Expect(err).NotTo(HaveOccurred())

			// Run initial init
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			_, err = cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be idempotent when run multiple times", func() {
			settingsPath := filepath.Join(testProjectDir, ".claude", "settings.local.json")

			// Get settings after first run
			firstRunData, err := os.ReadFile(settingsPath)
			Expect(err).NotTo(HaveOccurred())

			// Run init again
			cmd := exec.Command(specwareBinary, "init", testProjectDir, "-y")
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred())

			outputStr := string(output)
			Expect(outputStr).To(ContainSubstring("Specware permissions already configured"))

			// Get settings after second run
			secondRunData, err := os.ReadFile(settingsPath)
			Expect(err).NotTo(HaveOccurred())

			// Settings should be identical
			var firstSettings, secondSettings spec.ClaudeSettings
			err = json.Unmarshal(firstRunData, &firstSettings)
			Expect(err).NotTo(HaveOccurred())
			err = json.Unmarshal(secondRunData, &secondSettings)
			Expect(err).NotTo(HaveOccurred())

			Expect(len(firstSettings.Permissions.Allow)).To(Equal(len(secondSettings.Permissions.Allow)))
			for _, entry := range firstSettings.Permissions.Allow {
				Expect(secondSettings.Permissions.Allow).To(ContainElement(entry))
			}
		})
	})

	Describe("Flag variations", func() {
		var testProjectDir string

		BeforeEach(func() {
			testProjectDir = filepath.Join(tempDir, "test-project")
		})

		It("should accept both -y and --yes flags", func() {
			// Test short flag
			cmd := exec.Command(specwareBinary, "init", testProjectDir+"-short", "-y")
			err := cmd.Run()
			Expect(err).NotTo(HaveOccurred())

			// Test long flag
			cmd = exec.Command(specwareBinary, "init", testProjectDir+"-long", "--yes")
			err = cmd.Run()
			Expect(err).NotTo(HaveOccurred())

			// Verify both created project structure
			Expect(filepath.Join(testProjectDir+"-short", ".spec")).To(BeADirectory())
			Expect(filepath.Join(testProjectDir+"-long", ".spec")).To(BeADirectory())
		})
	})
})
