package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()

	// Database
	srv.store = &dbStore{}
	err := srv.store.Open()
	if err != nil {
		return err
	}
	defer srv.store.Close()

	srv.store.GetMovies()

	http.HandleFunc("/", srv.serveHTTP)
	log.Printf("Serving on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
