package main

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

// 两件事：parse & validate
// 修改 headers 或 payload 对应的 base64，可验证 parse（会失败）
// 修改 signature 对应的 base64，可验证 validate（会失败）

func TestValidate(t *testing.T) {
	token_str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDA4MTUwNzYsInN1YiI6InBweiJ9.8BcXrxWeg6KmpmIspCFqrTrJbv0bB6-2vFRilhG0Y2E"

	token, err := jwt.Parse(token_str, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t.Valid)            // false，因为此时还未提供密码
		fmt.Println(t.Header, t.Claims) // 因为 headers 和 paylout 是明文的，所以一开始就能 parse
		fmt.Println(t.Method)           // Method 是 headers 里的一个字段
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("omamimamiho"), nil
	})

	if err != nil {
		t.Error("error on parsing or validating", err)
	}

	fmt.Println(token.Valid) // true, 上个 err 为 nil 时，已经表示 validate 成功
	fmt.Println(token.Claims)
	fmt.Println(token.Claims.GetSubject())        // ppz
	fmt.Println(token.Claims.GetExpirationTime()) // nil, nil
}
