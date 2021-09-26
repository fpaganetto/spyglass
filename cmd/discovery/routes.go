package main

import "net/http"

func routes() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/", discovery)

	return handler
}
