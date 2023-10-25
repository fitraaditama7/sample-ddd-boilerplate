package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up")
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}
