package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"sabigo/config"
	"sabigo/logger"
	"strings"
)

type User struct {
	UserId int
	Token  string
}

type ErrorMessage struct {
	StatusCode int
	Message    string
}

type ContextKey int

const UserIDKey ContextKey = iota

func AuthMiddleware(next http.Handler) http.Handler {
	var user User
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug.Println(r.Header)
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		DB := config.ConnectDatabase()
		defer DB.Close()
		err := DB.QueryRow("SELECT user_id FROM tokens WHERE token=?", reqToken).Scan(&user.UserId)
		if err != nil {
			logger.Info.Println(err.Error())
			errorMessage := ErrorMessage{
				StatusCode: 2,
				Message:    "You are not authorized to access this resource",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorMessage)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, user.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
