package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL Shortener API")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server Error: ", err)
	}
}
