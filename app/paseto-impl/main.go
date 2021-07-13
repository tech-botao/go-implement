package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/k0kubun/pp"
	"github.com/o1egl/paseto"
	"time"
)

func genPrivateKey(seed string) []byte {
	b, _ := hex.DecodeString(seed)
	privateKey := ed25519.PrivateKey(b)
	return privateKey
}

func localMode() {
	//key := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	key := genPrivateKey("Hello Paseto.")
	now := time.Now()
	exp := now.Add(1*time.Minute)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "Test",
		Issuer:     "test service",
		Jti:        "1234",
		Subject:    "test_subject",
		Expiration: exp,
		IssuedAt:   now,
		NotBefore:  nbt,
	}

	jsonToken.Set("data", "this is a signed message")
	footer := "some footer"
	v2 := paseto.NewV2()
	token, err := v2.Sign(key, jsonToken, footer)
	if err != nil {
		panic(err)
	}
	pp.Println(token)
}

func main() {
	localMode()

}
