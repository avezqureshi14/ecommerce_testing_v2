package main

import (
	"encoding/json"
	"net/http"
)

func handleHealthJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":        true,
		"accepted":  acceptedCount.Load(),
		"rejected":  rejectedCount.Load(),
		"pageviews": pageViews.Len(),
		"errors":    errorEvents.Len(),
		"vitals":    vitals.Len(),
	})
}
