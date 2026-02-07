package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/musab/portfolio-backend/data"
)

func ContactInfo(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Contact)
	}
}

func ContactSubmit(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var msg data.ContactMessage
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
			return
		}

		if msg.Name == "" || msg.Email == "" || msg.Message == "" {
			http.Error(w, `{"error":"name, email, and message are required"}`, http.StatusBadRequest)
			return
		}

		store.AddMessage(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Message received",
		})
	}
}
