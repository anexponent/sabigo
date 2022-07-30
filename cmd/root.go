package cmd

import (
	"os"
	"sabigo/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sabigo",
	Short: "Sabigo framework command",
	Long:  `The Sabigo golang framework root command`,
}

func Execute() {
	logger.Init()
	if err := rootCmd.Execute(); err != nil {
		logger.Error.Fatal(err)
		os.Exit(1)
	}
}
