package cmd

import (
	"log"
	"os"
	"sabigo/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sabigo",
	Short: "Sabigo framework command",
	Long:  `The Sabigo golang framework root command`,
}

func Execute() {
	utils.Log()
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
