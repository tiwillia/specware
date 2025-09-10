package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Feature specification commands",
}

var newRequirementsCmd = &cobra.Command{
	Use:   "new-requirements <short-name>",
	Short: "Create new feature specification directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
	},
}

var newImplementationPlanCmd = &cobra.Command{
	Use:   "new-implementation-plan <short-name>",
	Short: "Create implementation plan for existing feature",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
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