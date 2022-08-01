package utils

import (
	"crypto/rand"
	"sabigo/logger"
)

func GenerateToken(len int) (string, error) {
	key := make([]byte, len)
	token := string(key[:])
	_, err := rand.Read(key)
	if err != nil {
		logger.Info.Println(err)
		return token, err
	}
	token = string(key[:])
	return token, nil
}
