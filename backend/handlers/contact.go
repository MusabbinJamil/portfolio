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

func MessagesGet(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectedKey := os.Getenv("ANALYTICS_KEY")
		if expectedKey == "" {
			http.Error(w, `{"error":"endpoint not configured"}`, http.StatusForbidden)
			return
		}

		providedKey := r.URL.Query().Get("key")
		if providedKey != expectedKey {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusForbidden)
			return
		}

		messages := store.GetMessages()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"total":    len(messages),
			"messages": messages,
		})
	}
}

func LoadMessagesFromCSV(store *data.Store) error {
	filePath := "messages.csv"
	f, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		if len(record) < 4 {
			continue
		}
		store.AddMessage(data.ContactMessage{
			Timestamp: record[0],
			Name:      record[1],
			Email:     record[2],
			Message:   record[3],
		})
	}
	return nil
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

		msg.Timestamp = time.Now().UTC().Format(time.RFC3339)
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
		msg.Timestamp,
		msg.Name,
		msg.Email,
		msg.Message,
	})
}
