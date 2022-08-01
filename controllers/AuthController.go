package controllers

import (
	"encoding/json"
	"net/http"
	"sabigo/config"
	"sabigo/logger"
	"sabigo/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logger.Init()

	DB := config.ConnectDatabase()
	defer DB.Close()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Info.Println("Error: " + err.Error())
		// http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)
	insert := user.InsertUser()
	logger.Debug.Println(insert)
	logger.Debug.Println(user)
	// statement, err := DB.Prepare("insert into users(username, email, phone, password, status, created_at) values (?, ?, ?, ?, ?, ?)")
	// if err != nil {
	// 	logger.Error.Println(err.Error())
	// }
	// //execute
	// result, err := statement.Exec(reqData.Username, reqData.Email, reqData.Phone, reqData.Password, 1, time.Now())
	// if err != nil {
	// 	logger.Error.Println(err.Error())
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	logger.Error.Println(err.Error())
	// }
	// fmt.Println("Insert id", id)

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
