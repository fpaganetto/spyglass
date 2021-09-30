package main

import (
	"fmt"
	"net/http"
)

func routes() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/", discovery)
	handler.Handle("/auth", verifyToken(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Authentication passed")
	})))

	return handler
}
