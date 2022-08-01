package models

import (
	"os"
	"sabigo/config"
	"sabigo/logger"
	"sabigo/utils"
	"time"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Token struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

func (u User) InsertUser() Token {

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
	token_key, err := utils.GenerateToken(128)
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	token := Token{
		UserId: int(user_id),
		Token:  token_key,
	}
	_, err = stmt.Exec(token.UserId, token.Token, time.Now())
	if err != nil {
		tx.Rollback()
		logger.Error.Println(err)
		os.Exit(1)
	}
	tx.Commit()
	return token
}
