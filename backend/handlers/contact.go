package handlers

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"time"

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

		if err := saveToCSV(msg); err != nil {
			http.Error(w, `{"error":"failed to save message"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Message received",
		})
	}
}

func saveToCSV(msg data.ContactMessage) error {
	filePath := "messages.csv"
	isNew := false

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		isNew = true
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if isNew {
		if err := w.Write([]string{"timestamp", "name", "email", "message"}); err != nil {
			return err
		}
	}

	return w.Write([]string{
		time.Now().Format(time.RFC3339),
		msg.Name,
		msg.Email,
		msg.Message,
	})
}
