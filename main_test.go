package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type payload struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
}

func TestMain(t *testing.T) {
	header := header{
		"HS256",
		"JWT",
	}
	header_str := json_str(header)
	header_base64url := base64url(header_str)
	fmt.Println(header_str, header_base64url)

	payload := payload{
		"ppz's server(the macmini)",
		"user is ppz",
	}
	payload_str := json_str(payload)
	payload_base64url := base64url(payload_str)
	fmt.Println(payload_str, payload_base64url)

	signature := sign(header_base64url, payload_base64url, "omamimamiho")
	signature_base64url := base64url(signature)
	fmt.Println(
		signature_base64url ==
			"0exxSO6DjWMyFl7tFrFar2pfrKtClbb6a4w8_rHSZqw", // 这个字符串来自 main
	)
}

func json_str(obj any) string {
	result, _ := json.Marshal(obj)
	return string(result)
}

func base64url(input string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(input))
}

func sign(header, payload, secret string) string {
	// https://pkg.go.dev/crypto/hmac
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(header + "." + payload))
	return string(mac.Sum(nil))
}
