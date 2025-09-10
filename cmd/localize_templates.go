package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"specware/internal/spec"
)

var localizeTemplatesCmd = &cobra.Command{
	Use:   "localize-templates",
	Short: "Create project-specific templates",
	Long: `Copies embedded templates to .spec/templates/ directory for project-specific customization.

This allows you to modify templates locally for your project without affecting the embedded defaults.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		if err := spec.LocalizeTemplates(cwd); err != nil {
			fmt.Printf("Error localizing templates: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Templates localized to .spec/templates/")
	},
}