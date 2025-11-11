package app

import (
	"os"

	"github.com/Tortik3000/gomodinfo/internal/messages"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomodinfo",
	Short: "CLI for analysis go.mod and dependencies",
	Long:  messages.RootCmdLongInfo,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
