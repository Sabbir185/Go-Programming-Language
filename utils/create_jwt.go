package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := Base64UrlEncode(byteArrayHeader)

	byteArrayData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	dataB64 := Base64UrlEncode(byteArrayData)

	message := headerB64 + "." + dataB64

	byteArraySecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := Base64UrlEncode(signature)

	jwt := headerB64 + "." + dataB64 + "." + signatureB64

	return jwt, nil
}

func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}


/*

	jwt, err := utils.CreateJWT("my-secret", utils.Payload{
		Sub:         1,
		FirstName:   "Sabbir",
		LastName:    "Ahmmed",
		Email:       "sabbir.cse.nub@gmail.com",
		IsShopOwner: true,
	})
	if err != nil {
		panic(err)
	}
	println(jwt)

*/