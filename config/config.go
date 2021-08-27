package config

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
)

func LoadEnvironmentalVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error Loading .env file..", err)
	}

	return os.Getenv(key)
}

func ConnectDatabase() *sql.DB {
	fmt.Println("Checking the database parameters")

	//Get Database Driver used from the environment
	driver := LoadEnvironmentalVariables("DB_DRIVER")
	host := LoadEnvironmentalVariables("DB_HOST")
	user := LoadEnvironmentalVariables("DB_USER")
	password := LoadEnvironmentalVariables("DB_PASSWORD")
	port := LoadEnvironmentalVariables("DB_PORT")
	dbname := LoadEnvironmentalVariables("DB_NAME")
	ssl := LoadEnvironmentalVariables("SSL_MODE")

	var db *sql.DB
	var err error

	// fmt.Println("Configured Driver:", driver)

	// fmt.Println("Connecting to the database...")
	if driver == "mysql" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")/"+dbname)
		if err != nil {
			fmt.Println("Unable to commect to database", err.Error())
			os.Exit(1)
		}
	} else if driver == "psql" {
		connection_string := fmt.Sprintf("port=%s host=%s user=%s "+"password=%s dbname=%s sslmode=%s",
			port, host, user, password, dbname, ssl)
		db, err = sql.Open("postgres", connection_string)
		if err != nil {
			fmt.Println("Error Connecting to Database", err.Error())
			os.Exit(1)
		}

	} else if driver == "sqlite" {
		db, err = sql.Open("sqlite3", dbname)
		if err != nil {
			fmt.Printf("Error opening database: %v", err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println("Database Driver:", driver, " not supported yet")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error Connecting to the database", err.Error())
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to connect to database", err.Error())
		os.Exit(1)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error Pinging Database")
	}
	fmt.Println("Database connected!")

	return db
}
