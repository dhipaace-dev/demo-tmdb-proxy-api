package handlers

import (
	"io"
	"net/http"
	//"os"
)

const targetServer = "https://api.themoviedb.org"
const apiKey = "bf5c57d57d1c81706c6ef4794e8d753e"

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	//apiKey := os.Getenv("TMDB_API_KEY")

	targetURL := targetServer + r.URL.Path

	// Preserve query parameters
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	req, err := http.NewRequest(
		r.Method,
		targetURL,
		r.Body,
	)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy request headers
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Add TMDB API key
	query := req.URL.Query()
	query.Set("api_key", apiKey)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Target server unavailable", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, _ = io.Copy(w, resp.Body)
}

