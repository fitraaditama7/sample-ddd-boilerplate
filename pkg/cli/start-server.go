package cli

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/http"
	"ddd-boilerplate/pkg/logger"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "start http server",
	Long:  "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		logger.InitializeLogger()

		http.StartServer(cfg)
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
