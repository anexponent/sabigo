package utils

import (
	"net/http"
	"sabigo/logger"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func Validate(w http.ResponseWriter, r *http.Request, model interface{}) bool {
	validate = validator.New()
	err := validate.Struct(model)
	hasError := HasError(w, err, "Error: ", http.StatusBadRequest)
	if hasError {
		logger.Error.Println("Validation Failed")
		return false
	} else {
		logger.Info.Println("Validation Successful")
		return true
	}
}
