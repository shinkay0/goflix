package main

import (
	"log"
	"net/http"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
