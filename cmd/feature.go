package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tiwillia/specware/internal/spec"
)

var featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Feature specification commands",
}

var newRequirementsCmd = &cobra.Command{
	Use:   "new-requirements <short-name>",
	Short: "Create new feature specification directory",
	Long: `Creates a new feature specification directory with sequential numbering.

This command creates a directory with the pattern XXX-<short-name> where XXX is a 
sequential number starting from 001. The directory will contain:
- requirements.md (copied from localized or embedded template)
- context-requirements.md (for tracking Q&A sessions and context gathering)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortName := args[0]
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		createdFiles, err := spec.CreateNewRequirements(cwd, shortName)
		if err != nil {
			fmt.Printf("Error creating feature requirements: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created feature requirements for '%s'\n", shortName)
		fmt.Println("\nCreated files:")
		for _, file := range createdFiles {
			fmt.Printf("  %s\n", file)
		}
	},
}

var newImplementationPlanCmd = &cobra.Command{
	Use:   "new-implementation-plan <short-name>",
	Short: "Create implementation plan for existing feature",
	Long: `Creates an implementation plan for an existing feature specification.

The feature directory must already exist (created with new-requirements). This command adds:
- implementation-plan.md (copied from localized or embedded template) 
- context-implementation-plan.md (for tracking implementation Q&A sessions and context gathering)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortName := args[0]
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		createdFiles, err := spec.CreateNewImplementationPlan(cwd, shortName)
		if err != nil {
			fmt.Printf("Error creating implementation plan: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created implementation plan for feature '%s'\n", shortName)
		fmt.Println("\nCreated files:")
		for _, file := range createdFiles {
			fmt.Printf("  %s\n", file)
		}
	},
}

var updateStateCmd = &cobra.Command{
	Use:   "update-state <short-name> <status>",
	Short: "Update the status of a feature specification",
	Long: `Updates the status of a feature specification in the .spec-status.json file.

The status value can be any string, but suggested values for use with the specify.md 
Claude command include:
- "Requirements Gathering"
- "Requirements Context Gathering"
- "Requirements Expert Q&A"
- "Requirements Complete"
- "Requirements Interactive Review"
- "Implementation Planning"
- "Implementation Plan Q&A"
- "Implementation Plan Generated"
- "Implementation Plan Interactive Review"
- "Implementation Planning Complete"`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		shortName := args[0]
		status := args[1]
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		if err := spec.UpdateFeatureStatus(cwd, shortName, status); err != nil {
			fmt.Printf("Error updating feature status: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Updated status for feature '%s' to '%s'\n", shortName, status)
	},
}

func init() {
	featureCmd.AddCommand(newRequirementsCmd)
	featureCmd.AddCommand(newImplementationPlanCmd)
	featureCmd.AddCommand(updateStateCmd)
}