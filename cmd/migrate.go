package cmd

import (
	"fmt"
	"sabigo/config"
	"sabigo/utils"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/rubenv/sql-migrate"
	"time"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "sabigo database migrations tool",
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Unable to read flag `name`", err.Error())
			color.Unset()
			return
		}
		//get migrations directory
		config_migration_dir := config.LoadEnvironmentalVariables("MIG_DIR")
		migrations_path := filepath.Join(".", config_migration_dir)
		err = os.MkdirAll(migrations_path, os.ModePerm)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
			color.Unset()
		}
		
		today := time.Now().Format("20060102150405")
		migration_file_name := migrations_path + "/" + today + "-" + name + ".sql"
		//check if file is existing 
		exists, err :=utils.Exists(migration_file_name)
		
		if exists ==true{
			color.Set(color.FgRed)
			fmt.Println("Error creating migration: " + name +" File Exists Already")
			os.Exit(1)
			color.Unset()
		}

		file, err := os.OpenFile("./"+ migration_file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println(err)
			os.Exit(1)
			color.Unset()
		}else{
			defer file.Close()
			_, err := file.WriteString("-- +migrate Up\n-- +migrate Down")
			if err != nil{
				color.Set(color.FgRed)
				_ = os.Remove(name + ".sql")
				fmt.Println("Error creating migration: " + name)
				os.Exit(1)
				color.Unset()
			}
		}
		color.Set(color.FgYellow)
		fmt.Println("Migration " + name + ".sql" + "created successfully")
		color.Unset()
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		config_migration_dir := config.LoadEnvironmentalVariables("MIG_DIR")
		migrations_path := filepath.Join(".", config_migration_dir)

		_, err := cmd.Flags().GetInt("step")
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Unable to read flag `step`")
			color.Unset()
			return
		}
		DB := config.ConnectDatabase()
		migrations := &migrate.FileMigrationSource{
			Dir: migrations_path,
		}
		n, err := migrate.Exec(DB, "mysql", migrations, migrate.Up)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Error occcured:", err)
			os.Exit(1)
			color.Unset()
		}
		if n == 0 {
			color.Set(color.FgYellow)
			fmt.Println("No migration to apply")
			color.Unset()
			return
		}
		color.Set(color.FgGreen)
		fmt.Printf("Applied %d migrations!\n", n)
		color.Unset()
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		config_migration_dir := config.LoadEnvironmentalVariables("MIG_DIR")
		migrations_path := filepath.Join(".", config_migration_dir)

		_, err := cmd.Flags().GetInt("step")
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Unable to read flag `step`")
			color.Unset()
			return
		}
		DB := config.ConnectDatabase()
		migrations := &migrate.FileMigrationSource{
			Dir: migrations_path,
		}
		n, err := migrate.Exec(DB, "mysql", migrations, migrate.Down)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Error occcured:", err)
			os.Exit(1)
			color.Unset()
		}
		if n == 0 {
			color.Set(color.FgYellow)
			fmt.Println("No migration to apply")
			color.Unset()
			return
		}
		color.Set(color.FgGreen)
		fmt.Printf("Applied %d migrations!\n", n)
		color.Unset()
	},
}

func init() {
	// Add "--name" flag to "create" command
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")
	// Add "--step" flag to both "up" and "down" command
	migrateUpCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	migrateDownCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	// Add "create", "up" and "down" commands to the "migrate" command
	migrateCmd.AddCommand(migrateUpCmd, migrateCreateCmd, migrateDownCmd)
	// Add "migrate" command to the root command
	rootCmd.AddCommand(migrateCmd)
}
