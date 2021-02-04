package cmds

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "demo-client",
	Short: "A command line tool to interact with the bank demo.",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
