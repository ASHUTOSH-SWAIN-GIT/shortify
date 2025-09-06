package internals

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shortify/internals/db"
	"shortify/internals/db/models"
	"time"
)

type ShortenResponse struct {
	ShortURL  string `json:"short_url"`
	LongURL   string `json:"long_url"`
	ShortCode string `json:"short_code"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Database connection
	dataSourceName := "root:your_new_password@tcp(localhost:3306)/url_shortner?parseTime=true"
	fmt.Printf("Attempting to connect to database with: %s\n", dataSourceName)

	store, err := db.NewMySqlStore(dataSourceName)
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("Database connected successfully!")
	defer store.DB.Close()

	// Generate short code
	shortCode := GenerateShortCode()

	// Check if short code already exists and regenerate if needed
	ctx := context.Background()
	for {
		exists, err := store.Exists(ctx, shortCode)
		if err != nil {
			fmt.Printf("Error checking if short code exists: %v\n", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		if !exists {
			break
		}
		shortCode = GenerateShortCode() // Regenerate if exists
	}

	// Create URL model
	url := &models.URL{
		ShortCode: shortCode,
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}

	// Save to database
	if err := store.Save(ctx, url); err != nil {
		fmt.Printf("Error saving URL to database: %v\n", err)
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := ShortenResponse{
		ShortURL:  fmt.Sprintf("http://localhost:8080/s/%s", shortCode),
		LongURL:   longURL,
		ShortCode: shortCode,
	}

	// Set JSON content type and send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Successfully shortened URL: %s -> %s\n", longURL, response.ShortURL)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract short code from URL path
	shortCode := r.URL.Path[len("/s/"):]
	if shortCode == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	// Database connection
	dataSourceName := "root:your_new_password@tcp(localhost:3306)/url_shortner?parseTime=true"
	store, err := db.NewMySqlStore(dataSourceName)
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer store.DB.Close()

	// Get URL from database
	ctx := context.Background()
	url, err := store.Get(ctx, shortCode)
	if err != nil {
		if err == db.ErrNotFound {
			http.Error(w, "Short URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Redirect to the long URL
	http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
}
