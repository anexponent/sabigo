package controllers

import (
	"encoding/json"
	"net/http"
	"sabigo/logger"
	"sabigo/middleware"
	"sabigo/models"
	"sabigo/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logger.Init()

	var user models.User

	//decode json into the body
	err := json.NewDecoder(r.Body).Decode(&user)
	if utils.HasError(w, err, "Decoding Eror: ", http.StatusBadRequest) {
		return
	}

	//validate user struct
	validated := utils.Validate(w, r, user)
	if !validated {
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
	logger.Init()

	var userJson models.UserLoginRequest
	var user models.UserSelected

	err := json.NewDecoder(r.Body).Decode(&userJson)
	if utils.HasError(w, err, "Decoding Eror: ", http.StatusBadRequest) {
		return
	}

	//validate the json
	if !utils.Validate(w, r, userJson) {
		return
	}

	login := user.SelectUser(userJson, w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(login)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value("user_id").(int)
	type logout struct {
		StatusCode int
		Message    string
	}
	_logout := logout{
		StatusCode: 1,
		Message:    "Logout Error",
	}
	userId := r.Context().Value(middleware.UserIDKey).(int)
	if models.Logout(userId) {
		_logout = logout{
			StatusCode: 0,
			Message:    "Logout Successfully",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(_logout)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(_logout)
}

func Profile(w http.ResponseWriter, r *http.Request) {

}

func Reset(w http.ResponseWriter, r *http.Request) {

}

func Flush(w http.ResponseWriter, r *http.Request) {

}
