package models

import (
	"os"
	"sabigo/config"
	"sabigo/logger"
	"time"

	"github.com/thanhpk/randstr"
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
	//insert into
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
