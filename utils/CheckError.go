package utils

import (
	"encoding/json"
	"net/http"
	"sabigo/logger"
)

type ErrorMessage struct {
	Message string `json:"error"`
}

func HasError(w http.ResponseWriter, err error, message string, status int) bool {
	if err != nil {
		logger.Info.Println(message + err.Error())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(ErrorMessage{Message: err.Error()})
		return true
	} else {
		return false
	}
}
