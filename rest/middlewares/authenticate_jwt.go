package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", 401)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		cnf := config.GetConfig()

		byteArraySecret := []byte(cnf.JwtSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArraySecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := utils.Base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized Access", 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
