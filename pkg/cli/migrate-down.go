package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down")
		},
	}

	migrateCmd.AddCommand(migrateDownCmd)
}
