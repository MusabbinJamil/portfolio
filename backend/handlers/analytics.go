package handlers

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/musab/portfolio-backend/data"
)

func getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return strings.TrimSpace(ip)
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func hashIP(ip string) string {
	h := sha256.Sum256([]byte(ip))
	return hex.EncodeToString(h[:])
}

// rate limiter

type rateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	rl := &rateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			rl.cleanup()
		}
	}()
	return rl
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	times := rl.requests[ip]
	valid := times[:0]
	for _, t := range times {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}

	if len(valid) >= rl.limit {
		rl.requests[ip] = valid
		return false
	}

	rl.requests[ip] = append(valid, now)
	return true
}

func (rl *rateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cutoff := time.Now().Add(-rl.window)
	for ip, times := range rl.requests {
		valid := times[:0]
		for _, t := range times {
			if t.After(cutoff) {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(rl.requests, ip)
		} else {
			rl.requests[ip] = valid
		}
	}
}

var analyticsLimiter = newRateLimiter(60, time.Minute)

type analyticsRequest struct {
	EventType string `json:"eventType"`
	Target    string `json:"target"`
	Label     string `json:"label"`
	Referrer  string `json:"referrer"`
}

func AnalyticsTrack(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)

		if !analyticsLimiter.allow(clientIP) {
			http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
			return
		}

		var req analyticsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
			return
		}

		if req.EventType != "pageview" && req.EventType != "click" {
			http.Error(w, `{"error":"eventType must be pageview or click"}`, http.StatusBadRequest)
			return
		}

		if req.EventType == "click" && req.Target == "" {
			http.Error(w, `{"error":"target is required for click events"}`, http.StatusBadRequest)
			return
		}

		if len(req.Target) > 100 {
			req.Target = req.Target[:100]
		}
		if len(req.Label) > 200 {
			req.Label = req.Label[:200]
		}
		if len(req.Referrer) > 500 {
			req.Referrer = req.Referrer[:500]
		}

		event := data.AnalyticsEvent{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			IP:        hashIP(clientIP),
			UserAgent: r.Header.Get("User-Agent"),
			EventType: req.EventType,
			Target:    req.Target,
			Label:     req.Label,
			Referrer:  req.Referrer,
		}

		store.AddEvent(event)

		if err := saveAnalyticsToCSV(event); err != nil {
			log.Printf("Failed to save analytics event: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}
}

func AnalyticsGet(store *data.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectedKey := os.Getenv("ANALYTICS_KEY")
		if expectedKey == "" {
			http.Error(w, `{"error":"analytics endpoint not configured"}`, http.StatusForbidden)
			return
		}

		providedKey := r.URL.Query().Get("key")
		if providedKey != expectedKey {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusForbidden)
			return
		}

		events := store.GetEvents()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"total":  len(events),
			"events": events,
		})
	}
}

func saveAnalyticsToCSV(event data.AnalyticsEvent) error {
	filePath := "analytics.csv"
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
		if err := w.Write([]string{"timestamp", "ip", "user_agent", "event_type", "target", "label", "referrer"}); err != nil {
			return err
		}
	}

	return w.Write([]string{
		event.Timestamp,
		event.IP,
		event.UserAgent,
		event.EventType,
		event.Target,
		event.Label,
		event.Referrer,
	})
}

func LoadAnalyticsFromCSV(store *data.Store) error {
	filePath := "analytics.csv"
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
		if len(record) < 7 {
			continue
		}
		store.AddEvent(data.AnalyticsEvent{
			Timestamp: record[0],
			IP:        record[1],
			UserAgent: record[2],
			EventType: record[3],
			Target:    record[4],
			Label:     record[5],
			Referrer:  record[6],
		})
	}
	return nil
}
