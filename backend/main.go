package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/musab/portfolio-backend/data"
	"github.com/musab/portfolio-backend/handlers"
)

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func main() {
	store := data.NewStore()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hero", cors(handlers.Hero(store)))
	mux.HandleFunc("GET /api/about", cors(handlers.About(store)))
	mux.HandleFunc("GET /api/projects", cors(handlers.Projects(store)))
	mux.HandleFunc("GET /api/contact", cors(handlers.ContactInfo(store)))
	mux.HandleFunc("POST /api/contact", cors(handlers.ContactSubmit(store)))

	mux.HandleFunc("GET /api/health", cors(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok"}`)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Backend listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
