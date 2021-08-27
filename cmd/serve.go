package cmd

import (
	"log"
	"sabigo/server"

	"github.com/spf13/cobra"
)

var serveCmd *cobra.Command

func init() {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve Command to start sabiserver. Use command with port argument. e.g 'sabigo serve 9000'",
		Long:  `Serve command to start web sabiserver`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Server Started through sabigo command")
			server.Serve()
		},
	}

	rootCmd.AddCommand(serveCmd)
}
