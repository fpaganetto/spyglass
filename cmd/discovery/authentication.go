package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func verifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verifying token")
		token := r.Header.Get("Authorization")

		if token == "" {
			forbidden(w, r)
			return
		}

		token = strings.Replace(token, "Bearer ", "", -1)
		x5c := "YOUR_X5C_KEY"
		mySigningKey := []byte(fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", x5c))
		parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(mySigningKey)

		if err != nil {
			fmt.Println(err.Error())
			forbidden(w, r)
			return
		}

		myToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}

			return parsedKey, nil
		})

		if err != nil {
			w.Write([]byte(err.Error()))
			forbidden(w, r)
			return
		}

		fmt.Print(myToken)

		next.ServeHTTP(w, r)
	})
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
