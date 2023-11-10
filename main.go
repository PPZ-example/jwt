package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "ppz's server(the macmini)",
			"sub": "user is ppz",
		},
	)
	token_str, err := token.SignedString([]byte("omamimamiho"))
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwcHoncyBzZXJ2ZXIodGhlIG1hY21pbmkpIiwic3ViIjoidXNlciBpcyBwcHoifQ.0exxSO6DjWMyFl7tFrFar2pfrKtClbb6a4w8_rHSZqw

	fmt.Println(token_str, err)
}
