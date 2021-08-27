package cmd

import (
	"sabigo/config"

	"github.com/spf13/cobra"
)

var dbstatusCmd *cobra.Command

func init() {
	dbstatusCmd = &cobra.Command{
		Use:   "dbstatus",
		Short: "Check the status of the database configured for the your project",
		Long:  `Check the status of the database configured for the your project`,
		Run: func(cmd *cobra.Command, args []string) {
			config.ConnectDatabase()
		},
	}

	rootCmd.AddCommand(dbstatusCmd)
}
