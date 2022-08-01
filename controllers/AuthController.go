package controllers

import (
	"encoding/json"
	"net/http"
	"sabigo/config"
	"sabigo/logger"
	"sabigo/models"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func Register(w http.ResponseWriter, r *http.Request) {

	logger.Init()

	DB := config.ConnectDatabase()
	defer DB.Close()

	var user models.User

	//decode json into the body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Info.Println("Error: " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	//validate user struct
	validate = validator.New()
	err = validate.Struct(user)
	// var errors []string
	if err != nil {
		logger.Debug.Println(("Validation Error: " + err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)
	insert := user.InsertUser()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(insert)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func Profile(w http.ResponseWriter, r *http.Request) {

}

func Reset(w http.ResponseWriter, r *http.Request) {

}

func Flush(w http.ResponseWriter, r *http.Request) {

}
