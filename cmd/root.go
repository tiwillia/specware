package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "specware",
	Short: "Spec-driven workflow enablement tool",
	Long:  `A tool to facilitate spec-driven development workflows through Claude Code AI Coding Assistant.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(localizeTemplatesCmd)
	rootCmd.AddCommand(featureCmd)
}