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
		
		fmt.Println("\nNext steps:")
		fmt.Println("  1. Use 'claude' and '/specify' to begin feature specification")
		fmt.Println("  2. Optionally run 'specware localize-templates' for custom templates")
	},
}