package main

import (
	"encoding/json"
	"net/http"
)

var vitals = newBatcher(200, 0)

var knownVitals = map[string]bool{"LCP": true, "CLS": true, "INP": true, "FCP": true, "TTFB": true}

func handleVitals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var v VitalBeacon
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 32*1024)).Decode(&v); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if !knownVitals[v.Name] {
		http.Error(w, "unknown vital name", http.StatusBadRequest)
		return
	}
	if err := checkURL(v.PageURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if vitals.Add(v) {
		_ = vitals.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}
