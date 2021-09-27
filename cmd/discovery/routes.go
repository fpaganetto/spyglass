package main

import "net/http"

func routes() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/", discovery)
	handler.Handle("/auth", verifyToken(http.HandlerFunc(discovery)))

	return handler
}
