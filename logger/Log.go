package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sabigo/config"
	"time"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
)

func Init() {
	logChannel := config.LoadEnvironmentalVariables("LOG_CHANNEL")
	environment := config.LoadEnvironmentalVariables("ENVIRONMENT")
	logPath := filepath.Join(".", "storage/logs")
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		log.Fatal(err.Error())
		// os.Exit(1)
	}

	var logFileName string

	today := time.Now().Format("01-02-2006")

	if logChannel == "daily" {
		logFileName = logPath + "/sabi-" + today + "-" + environment + ".log"
	} else {
		logFileName = logPath + "/sabi" + "-" + environment + ".log"
	}

	file, err := os.OpenFile("./"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//log.SetOutput(file)
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(file, "DEBUG", log.Ldate|log.Ltime|log.Lshortfile)
}
