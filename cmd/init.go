package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"specware/internal/spec"
)

var initCmd = &cobra.Command{
	Use:   "init <directory>",
	Short: "Initialize project to support spec-driven-workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := args[0]
		
		if err := spec.InitProject(targetDir); err != nil {
			fmt.Printf("Error initializing project: %v\n", err)
			return
		}
		
		fmt.Printf("Successfully initialized spec-driven workflow in %s\n", targetDir)
		fmt.Println("Next steps:")
		fmt.Println("  1. Use 'claude' and '/specify' to begin feature specification")
		fmt.Println("  2. Optionally run 'specware localize-templates' for custom templates")
	},
}