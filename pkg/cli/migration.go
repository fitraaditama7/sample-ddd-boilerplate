package cli

import "github.com/spf13/cobra"

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate cmd",
	Long:  "migrate command",
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
