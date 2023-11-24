package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerate(t *testing.T) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": "ppz",
			"iat": time.Now().Unix(),
		},
	)

	token_str, err := token.SignedString([]byte("omamimamiho"))
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(token_str)
		// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDA4MTUwNzYsInN1YiI6InBweiJ9.8BcXrxWeg6KmpmIspCFqrTrJbv0bB6-2vFRilhG0Y2E
	}
}
