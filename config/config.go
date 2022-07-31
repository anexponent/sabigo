package config

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/joho/godotenv"
	"github.com/fatih/color"
)

func LoadEnvironmentalVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		color.Set(color.FgRed)
		log.Println("Error Loading .env file..", err)
		color.Unset()
	}

	return os.Getenv(key)
}

func ConnectDatabase() *sql.DB {
	color.Set(color.FgYellow)
	fmt.Println("Checking the database parameters")
	color.Unset()
	//Get Database Driver used from the environment
	driver := LoadEnvironmentalVariables("DB_DRIVER")
	host := LoadEnvironmentalVariables("DB_HOST")
	user := LoadEnvironmentalVariables("DB_USER")
	password := LoadEnvironmentalVariables("DB_PASSWORD")
	port := LoadEnvironmentalVariables("DB_PORT")
	dbname := LoadEnvironmentalVariables("DB_NAME")
	ssl := LoadEnvironmentalVariables("SSL_MODE")
	parseTime := LoadEnvironmentalVariables("PARSE_TIME")

	var db *sql.DB
	var err error

	// fmt.Println("Configured Driver:", driver)

	// fmt.Println("Connecting to the database...")
	if driver == "mysql" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")/"+dbname + "?parseTime="+parseTime)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Unable to commect to database", err.Error())
			os.Exit(1)
			color.Unset()
		}
	} else if driver == "psql" {
		connection_string := fmt.Sprintf("port=%s host=%s user=%s "+"password=%s dbname=%s sslmode=%s",
			port, host, user, password, dbname, ssl)
		db, err = sql.Open("postgres", connection_string)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Error Connecting to Database", err.Error())
			os.Exit(1)
			color.Unset()
		}

	} else if driver == "sqlite" {
		db, err = sql.Open("sqlite3", dbname)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Printf("Error opening database: %v", err.Error())
			os.Exit(1)
			color.Unset()
		}
	} else {
		color.Set(color.FgRed)
		fmt.Println("Database Driver:", driver, " not supported yet")
		os.Exit(1)
		color.Unset()
	}
	if err != nil {
		color.Set(color.FgRed)
		fmt.Println("Error Connecting to the database", err.Error())
		os.Exit(1)
		color.Unset()
	}

	if err := db.Ping(); err != nil {
		color.Set(color.FgRed)
		fmt.Println("Unable to connect to database", err.Error())
		os.Exit(1)
		color.Unset()
	}

	err = db.Ping()

	if err != nil {
		color.Set(color.FgRed)
		fmt.Println("Error Pinging Database")
		color.Unset()
	}
	color.Set(color.FgGreen)
	fmt.Println("Database connected!")
	color.Unset()

	// defer db.Close()
	return db
}
