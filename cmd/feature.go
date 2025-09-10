package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"specware/internal/spec"
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
- q&a-requirements.md (for tracking Q&A sessions)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortName := args[0]
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		if err := spec.CreateNewRequirements(cwd, shortName); err != nil {
			fmt.Printf("Error creating feature requirements: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created feature requirements for '%s'\n", shortName)
	},
}

var newImplementationPlanCmd = &cobra.Command{
	Use:   "new-implementation-plan <short-name>",
	Short: "Create implementation plan for existing feature",
	Long: `Creates an implementation plan for an existing feature specification.

The feature directory must already exist (created with new-requirements). This command adds:
- implementation-plan.md (copied from localized or embedded template) 
- q&a-implementation-plan.md (for tracking implementation Q&A sessions)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortName := args[0]
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		if err := spec.CreateNewImplementationPlan(cwd, shortName); err != nil {
			fmt.Printf("Error creating implementation plan: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created implementation plan for feature '%s'\n", shortName)
	},
}

var updateStateCmd = &cobra.Command{
	Use:   "update-state <short-name> <status>",
	Short: "Update the status of a feature specification",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
	},
}

func init() {
	featureCmd.AddCommand(newRequirementsCmd)
	featureCmd.AddCommand(newImplementationPlanCmd)
	featureCmd.AddCommand(updateStateCmd)
}