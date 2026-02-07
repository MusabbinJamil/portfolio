package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/musab/portfolio-backend/data"
)

func Projects(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Projects)
	}
}
