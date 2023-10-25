package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcli",
	Short: "Root Command",
	Long:  "Root Command",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
