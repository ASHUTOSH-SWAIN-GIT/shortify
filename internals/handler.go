package internals

import (
	"fmt"
	"net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method ", http.StatusMethodNotAllowed)
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received url to shorten %s\n", longURL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL received! Check your server console."))
}
