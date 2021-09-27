package main

import (
	"fmt"
	"net/http"
)

func verifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verifying token")
		token := r.Header.Get("Authorization")
		if token == "" {
			forbidden(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
