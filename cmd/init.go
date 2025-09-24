package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiwillia/specware/internal/spec"
)

var yesFlag bool

var initCmd = &cobra.Command{
	Use:   "init <directory>",
	Short: "Initialize project to support spec-driven-workflow",
	Long: `Initialize project to support spec-driven-workflow

This command creates the following directory structure:
  .claude/commands/     - Claude Code command files (includes /specify workflow)
  .claude/agents/       - Claude Code agent files for specialized workflows
  .spec/                - Feature specifications directory
  .spec/config.json     - Configuration for workflow question counts
  .spec/README.md       - Documentation for the spec workflow

Optional modifications (user will be prompted):
  .claude/settings.local.json - Updates project permissions to allow specware
                                commands without prompting (personal settings only)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := args[0]

		createdFiles, err := spec.InitProject(targetDir)
		if err != nil {
			fmt.Printf("Error initializing project: %v\n", err)
			return
		}

		fmt.Printf("Successfully initialized spec-driven workflow in %s\n", targetDir)
		fmt.Println("\nCreated files:")
		for _, file := range createdFiles {
			fmt.Printf("  %s\n", file)
		}

		// Update Claude Code settings if requested
		if err := spec.UpdateClaudeSettings(targetDir, yesFlag); err != nil {
			fmt.Printf("Warning: Failed to update Claude Code settings: %v\n", err)
		}

		fmt.Println("\nNext steps:")
		fmt.Println("  1. Use 'claude' and '/specify' to begin feature specification")
		fmt.Println("  2. Optionally run 'specware localize-templates' for custom templates")
	},
}

func init() {
	initCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "automatically answer yes to all prompts")
}
