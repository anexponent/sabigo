package models

import (
	"net/http"
	"os"
	"sabigo/config"
	"sabigo/logger"
	"time"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,len=11"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type Token struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

type CreateUserResponse struct {
	StatusCode int    `json:"statusCode"`
	UserId     int    `json:"userId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Token      string `json:"token"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Username      string `json:"username"`
	Token         string `json:"token"`
}

type UserSelected struct {
	UserId   int
	Username string
	Password string
	Token    string
}

func (u User) InsertUser() CreateUserResponse {
	DB := config.ConnectDatabase()
	defer DB.Close()
	tx, err := DB.Begin()
	if err != nil {
		logger.Error.Println("Error saving to database: ", err.Error())
		os.Exit(1)
	}
	defer tx.Rollback()

	//insert into users table
	stmt, err := tx.Prepare(`INSERT INTO users(username, email, phone, password, status, created_at) VALUES (?, ?, ?, ?, ?, ?);`)
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	defer stmt.Close()

	user, err := stmt.Exec(u.Username, u.Email, u.Phone, u.Password, 1, time.Now())
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	user_id, err := user.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	//insert into token
	stmt, err = tx.Prepare(`INSERT INTO tokens(user_id, token, created_at) VALUES (?, ?, ?);`)
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	defer stmt.Close()
	token := Token{
		UserId: int(user_id),
		Token:  randstr.Hex(32),
	}
	_, err = stmt.Exec(token.UserId, token.Token, time.Now())
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	tx.Commit()
	response := CreateUserResponse{
		StatusCode: 0,
		UserId:     int(user_id),
		Username:   u.Username,
		Email:      u.Email,
		Phone:      u.Phone,
		Token:      token.Token,
	}
	return response
}

func (ul UserSelected) SelectUser(u UserLoginRequest, w http.ResponseWriter) UserLoginResponse {
	DB := config.ConnectDatabase()
	defer DB.Close()

	response := UserLoginResponse{
		StatusCode:    1,
		StatusMessage: "Wrong Credentials",
	}
	err := DB.QueryRow("select id, password from users where phone = ? or email = ? or username = ?", u.Username, u.Username, u.Username).Scan(&ul.UserId, &ul.Password)
	// utils.HasError(w, err, "Error Querying Database", http.StatusNotFound)
	if err != nil {
		logger.Error.Println("Error saving to database: ", err.Error())
		return response
	}

	err = bcrypt.CompareHashAndPassword([]byte(ul.Password), []byte(u.Password))
	if err != nil {
		logger.Error.Println("Error retrieving user from database: ", err.Error())
		return response
	}

	stmt, err := DB.Prepare(`INSERT INTO tokens(user_id, token, created_at) VALUES (?, ?, ?);`)
	if err != nil {
		logger.Error.Println(err)
		return response
	}
	defer stmt.Close()

	token := randstr.Hex(32)

	_, err = stmt.Exec(ul.UserId, token, time.Now())
	if err != nil {
		logger.Error.Println(err)
		return response
	}

	response = UserLoginResponse{
		StatusCode:    0,
		StatusMessage: "Success",
		Username:      u.Username,
		Token:         token,
	}
	return response
}
