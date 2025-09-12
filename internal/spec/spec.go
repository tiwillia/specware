package spec

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/tiwillia/specware/assets"
)

// InitProject initializes a project with spec-driven workflow support
func InitProject(targetDir string) ([]string, error) {
	var createdFiles []string
	// Create .claude/commands directory
	claudeDir := filepath.Join(targetDir, ".claude", "commands")
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create .claude/commands directory: %w", err)
	}

	// Create .claude/agents directory
	agentsDir := filepath.Join(targetDir, ".claude", "agents")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create .claude/agents directory: %w", err)
	}

	// Create .spec directory
	specDir := filepath.Join(targetDir, ".spec")
	if err := os.MkdirAll(specDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create .spec directory: %w", err)
	}

	// Copy commands from embedded assets
	commandsFS := assets.GetCommandsFS()
	err := fs.WalkDir(commandsFS, "commands", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Read file from embedded FS
		content, err := fs.ReadFile(commandsFS, path)
		if err != nil {
			return err
		}

		// Write to target directory
		relPath := strings.TrimPrefix(path, "commands/")
		targetPath := filepath.Join(claudeDir, relPath)
		createdFiles = append(createdFiles, filepath.Join(".claude", "commands", relPath))
		return os.WriteFile(targetPath, content, 0644)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to copy commands: %w", err)
	}

	// Copy agents from embedded assets
	agentsFS := assets.GetAgentsFS()
	err = fs.WalkDir(agentsFS, "agents", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Read file from embedded FS
		content, err := fs.ReadFile(agentsFS, path)
		if err != nil {
			return err
		}

		// Write to target directory
		relPath := strings.TrimPrefix(path, "agents/")
		targetPath := filepath.Join(agentsDir, relPath)
		createdFiles = append(createdFiles, filepath.Join(".claude", "agents", relPath))
		return os.WriteFile(targetPath, content, 0644)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to copy agents: %w", err)
	}

	// Create .spec/README.md
	specReadmeContent, err := assets.GetSpecReadme()
	if err != nil {
		return nil, fmt.Errorf("failed to get spec README content: %w", err)
	}
	
	readmePath := filepath.Join(specDir, "README.md")
	createdFiles = append(createdFiles, ".spec/README.md")
	if err := os.WriteFile(readmePath, specReadmeContent, 0644); err != nil {
		return nil, fmt.Errorf("failed to create .spec/README.md: %w", err)
	}

	// Create example spec directory
	exampleDir := filepath.Join(specDir, "000-example-spec")
	if err := os.MkdirAll(exampleDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create example spec directory: %w", err)
	}

	// Create .spec-status.json file
	statusPath := filepath.Join(exampleDir, ".spec-status.json")
	createdFiles = append(createdFiles, ".spec/000-example-spec/.spec-status.json")
	statusData := FeatureStatus{
		CurrentStep: "Not Started",
	}
	jsonData, err := json.MarshalIndent(statusData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal status data: %w", err)
	}
	if err := os.WriteFile(statusPath, jsonData, 0644); err != nil {
		return nil, fmt.Errorf("failed to create .spec-status.json file: %w", err)
	}

	return createdFiles, nil
}

// LocalizeTemplates copies embedded templates to project .spec/templates directory
func LocalizeTemplates(targetDir string) ([]string, error) {
	var createdFiles []string
	templatesDir := filepath.Join(targetDir, ".spec", "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create templates directory: %w", err)
	}

	// Copy templates from embedded assets
	templatesFS := assets.GetTemplatesFS()
	err := fs.WalkDir(templatesFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Read file from embedded FS
		content, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}

		// Write to target directory
		relPath := strings.TrimPrefix(path, "templates/")
		targetPath := filepath.Join(templatesDir, relPath)
		
		// Check if file already exists
		if _, err := os.Stat(targetPath); err == nil {
			fmt.Printf("Warning: Template file %s already exists, overwriting\n", relPath)
		}
		
		createdFiles = append(createdFiles, filepath.Join(".spec", "templates", relPath))
		return os.WriteFile(targetPath, content, 0644)
	})

	if err != nil {
		return nil, err
	}
	return createdFiles, nil
}

// GetNextFeatureNumber returns the next available feature number
func GetNextFeatureNumber(specDir string) (int, error) {
	entries, err := os.ReadDir(specDir)
	if err != nil {
		return 1, nil // If directory doesn't exist, start with 1
	}

	maxNum := 0
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if len(name) >= 3 && name[3] == '-' {
				if num, err := strconv.Atoi(name[:3]); err == nil {
					if num > maxNum {
						maxNum = num
					}
				}
			}
		}
	}

	return maxNum + 1, nil
}

// ValidateFeatureName validates that a feature short name is valid
func ValidateFeatureName(shortName string) error {
	if shortName == "" {
		return fmt.Errorf("feature name cannot be empty")
	}
	
	// Check for valid characters (alphanumeric, hyphens, underscores)
	matched, err := regexp.MatchString("^[a-zA-Z0-9_-]+$", shortName)
	if err != nil {
		return fmt.Errorf("error validating feature name: %w", err)
	}
	if !matched {
		return fmt.Errorf("feature name can only contain letters, numbers, hyphens, and underscores")
	}
	
	if len(shortName) > 50 {
		return fmt.Errorf("feature name must be 50 characters or less")
	}
	
	return nil
}

// getTemplate returns template content, preferring localized over embedded
func getTemplate(targetDir, templateName string) ([]byte, error) {
	// Try localized template first
	localPath := filepath.Join(targetDir, ".spec", "templates", templateName)
	if content, err := os.ReadFile(localPath); err == nil {
		return content, nil
	}
	
	// Fall back to embedded template
	templatesFS := assets.GetTemplatesFS()
	return fs.ReadFile(templatesFS, filepath.Join("templates", templateName))
}

// CreateNewRequirements creates a new feature requirements specification
func CreateNewRequirements(targetDir, shortName string) ([]string, error) {
	var createdFiles []string
	if err := ValidateFeatureName(shortName); err != nil {
		return nil, err
	}
	
	specDir := filepath.Join(targetDir, ".spec")
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		return nil, fmt.Errorf(".spec directory not found. Run 'specware init' first")
	}
	
	// Get next feature number
	featureNum, err := GetNextFeatureNumber(specDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get next feature number: %w", err)
	}
	
	// Create feature directory
	featureName := fmt.Sprintf("%03d-%s", featureNum, shortName)
	featureDir := filepath.Join(specDir, featureName)
	if err := os.MkdirAll(featureDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create feature directory: %w", err)
	}
	
	// Copy requirements template
	requirementsContent, err := getTemplate(targetDir, "requirements.md")
	if err != nil {
		return nil, fmt.Errorf("failed to get requirements template: %w", err)
	}
	
	requirementsPath := filepath.Join(featureDir, "requirements.md")
	createdFiles = append(createdFiles, filepath.Join(".spec", featureName, "requirements.md"))
	if err := os.WriteFile(requirementsPath, requirementsContent, 0644); err != nil {
		return nil, fmt.Errorf("failed to create requirements.md: %w", err)
	}
	
	// Create context file from template
	contextTemplate, err := getTemplate(targetDir, "context.md")
	if err != nil {
		return nil, fmt.Errorf("failed to get context template: %w", err)
	}
	
	// Replace placeholder with appropriate title
	contextContent := strings.Replace(string(contextTemplate), "[Feature Name]", "Requirements", 1)
	contextPath := filepath.Join(featureDir, "context-requirements.md")
	createdFiles = append(createdFiles, filepath.Join(".spec", featureName, "context-requirements.md"))
	if err := os.WriteFile(contextPath, []byte(contextContent), 0644); err != nil {
		return nil, fmt.Errorf("failed to create context-requirements.md: %w", err)
	}
	
	// Create .spec-status.json file
	statusPath := filepath.Join(featureDir, ".spec-status.json")
	createdFiles = append(createdFiles, filepath.Join(".spec", featureName, ".spec-status.json"))
	statusData := FeatureStatus{
		CurrentStep: "requirements-gathering",
	}
	jsonData, err := json.MarshalIndent(statusData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal status data: %w", err)
	}
	if err := os.WriteFile(statusPath, jsonData, 0644); err != nil {
		return nil, fmt.Errorf("failed to create .spec-status.json: %w", err)
	}
	
	return createdFiles, nil
}

// CreateNewImplementationPlan creates an implementation plan for an existing feature
func CreateNewImplementationPlan(targetDir, shortName string) ([]string, error) {
	var createdFiles []string
	if err := ValidateFeatureName(shortName); err != nil {
		return nil, err
	}
	
	specDir := filepath.Join(targetDir, ".spec")
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		return nil, fmt.Errorf(".spec directory not found. Run 'specware init' first")
	}
	
	// Find the feature directory
	featureDir, err := findFeatureDirectory(specDir, shortName)
	if err != nil {
		return nil, err
	}
	
	// Check if implementation plan already exists
	planPath := filepath.Join(featureDir, "implementation-plan.md")
	if _, err := os.Stat(planPath); err == nil {
		return nil, fmt.Errorf("implementation plan already exists for feature %s", shortName)
	}
	
	// Copy implementation plan template
	planContent, err := getTemplate(targetDir, "implementation-plan.md")
	if err != nil {
		return nil, fmt.Errorf("failed to get implementation plan template: %w", err)
	}
	
	// Extract feature name from directory path for relative path
	featureName := filepath.Base(featureDir)
	createdFiles = append(createdFiles, filepath.Join(".spec", featureName, "implementation-plan.md"))
	if err := os.WriteFile(planPath, planContent, 0644); err != nil {
		return nil, fmt.Errorf("failed to create implementation-plan.md: %w", err)
	}
	
	// Create context file from template
	contextTemplate, err := getTemplate(targetDir, "context.md")
	if err != nil {
		return nil, fmt.Errorf("failed to get context template: %w", err)
	}
	
	// Replace placeholder with appropriate title
	contextContent := strings.Replace(string(contextTemplate), "[Feature Name]", "Implementation Plan", 1)
	contextPath := filepath.Join(featureDir, "context-implementation-plan.md")
	createdFiles = append(createdFiles, filepath.Join(".spec", featureName, "context-implementation-plan.md"))
	if err := os.WriteFile(contextPath, []byte(contextContent), 0644); err != nil {
		return nil, fmt.Errorf("failed to create context-implementation-plan.md: %w", err)
	}
	
	return createdFiles, nil
}

// findFeatureDirectory finds a feature directory by short name
func findFeatureDirectory(specDir, shortName string) (string, error) {
	entries, err := os.ReadDir(specDir)
	if err != nil {
		return "", fmt.Errorf("failed to read spec directory: %w", err)
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			// Check if directory matches pattern: XXX-<shortName>
			if len(name) >= 4 && name[3] == '-' {
				if name[4:] == shortName {
					return filepath.Join(specDir, name), nil
				}
			}
		}
	}
	
	return "", fmt.Errorf("feature directory not found for %s. Run 'specware feature new-requirements %s' first", shortName, shortName)
}

// FeatureStatus represents the status information stored in .spec-status.json
type FeatureStatus struct {
	CurrentStep string `json:"current-step"`
}

// UpdateFeatureStatus updates the status of a feature specification
func UpdateFeatureStatus(targetDir, shortName, status string) error {
	if err := ValidateFeatureName(shortName); err != nil {
		return err
	}
	
	specDir := filepath.Join(targetDir, ".spec")
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		return fmt.Errorf(".spec directory not found. Run 'specware init' first")
	}
	
	// Find the feature directory
	featureDir, err := findFeatureDirectory(specDir, shortName)
	if err != nil {
		return err
	}
	
	// Update status file
	statusPath := filepath.Join(featureDir, ".spec-status.json")
	statusData := FeatureStatus{
		CurrentStep: status,
	}
	
	jsonData, err := json.MarshalIndent(statusData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal status data: %w", err)
	}
	
	if err := os.WriteFile(statusPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write status file: %w", err)
	}
	
	return nil
}