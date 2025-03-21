package main

import (
	"fmt"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "ok queue_depth=%d accepted=%d rejected=%d poor_vitals=%d\n", pageViews.Len(), acceptedCount.Load(), rejectedCount.Load(), vitalPoor.Load())
}
