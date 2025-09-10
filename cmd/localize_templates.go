package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var localizeTemplatesCmd = &cobra.Command{
	Use:   "localize-templates",
	Short: "Create project-specific templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
	},
}