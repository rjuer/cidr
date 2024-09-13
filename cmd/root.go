package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:     "cidr",
	Short:   "Application to operate with CIDR blocks",
	Long:    `Handles basic operations related to CIDR blocks.`,
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
