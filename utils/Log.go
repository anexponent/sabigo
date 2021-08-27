package utils

import (
	"log"
	"os"
	"sabigo/config"
)

func Log() {
	logLocation := config.LoadEnvironmentalVariables("LOG")

	file, err := os.OpenFile(logLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
