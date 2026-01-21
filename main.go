package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Response structure define karte hain
type StatusResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}

func main() {
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		res := StatusResponse{
			Message:   "Golang API v2 is running smoothly!",
			Timestamp: time.Now(),
			Status:    "Success",
		}

		// Header set karo taaki browser ko pata chale ye JSON hai
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("Server v2 starting at :8081...")
	http.ListenAndServe(":8081", nil)
}