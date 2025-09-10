package spec

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"specware/internal/assets"
)

// InitProject initializes a project with spec-driven workflow support
func InitProject(targetDir string) error {
	// Create .claude/commands directory
	claudeDir := filepath.Join(targetDir, ".claude", "commands")
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("failed to create .claude/commands directory: %w", err)
	}

	// Create .spec directory
	specDir := filepath.Join(targetDir, ".spec")
	if err := os.MkdirAll(specDir, 0755); err != nil {
		return fmt.Errorf("failed to create .spec directory: %w", err)
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
		return os.WriteFile(targetPath, content, 0644)
	})
	if err != nil {
		return fmt.Errorf("failed to copy commands: %w", err)
	}

	// Create .spec/README.md
	specReadmeContent, err := assets.GetSpecReadme()
	if err != nil {
		return fmt.Errorf("failed to get spec README content: %w", err)
	}
	
	readmePath := filepath.Join(specDir, "README.md")
	if err := os.WriteFile(readmePath, specReadmeContent, 0644); err != nil {
		return fmt.Errorf("failed to create .spec/README.md: %w", err)
	}

	// Create example spec directory
	exampleDir := filepath.Join(specDir, "000-example-spec")
	if err := os.MkdirAll(exampleDir, 0755); err != nil {
		return fmt.Errorf("failed to create example spec directory: %w", err)
	}

	// Create .spec-status file
	statusPath := filepath.Join(exampleDir, ".spec-status")
	statusContent := "status: initialized\ncreated: example\nphase: example\n"
	if err := os.WriteFile(statusPath, []byte(statusContent), 0644); err != nil {
		return fmt.Errorf("failed to create .spec-status file: %w", err)
	}

	return nil
}

// LocalizeTemplates copies embedded templates to project .spec/templates directory
func LocalizeTemplates(targetDir string) error {
	templatesDir := filepath.Join(targetDir, ".spec", "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return fmt.Errorf("failed to create templates directory: %w", err)
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
		return os.WriteFile(targetPath, content, 0644)
	})

	return err
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