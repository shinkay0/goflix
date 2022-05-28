package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  Store
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
