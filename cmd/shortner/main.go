package main

import (
	"fmt"
	"log"
	"net/http"

	"shortify/internals"
)

func main() {
	fmt.Println("Starting the URL shortner on port 8080")

	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", internals.ShortenHandler)

	fileserver := http.FileServer(http.Dir("./web/"))
	mux.Handle("/", fileserver)

	err := http.ListenAndServe(":8080", logRequest(mux))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
