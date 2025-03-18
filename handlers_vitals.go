package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var vitals = newBatcher(200, 0)

var knownVitals = map[string]bool{"LCP": true, "CLS": true, "INP": true, "FCP": true, "TTFB": true}

func handleVitals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !wantsJSON(r) {
		noteRejected()
		http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	if vitals.Len() >= 5000 {
		noteRejected()
		http.Error(w, "queue full", http.StatusServiceUnavailable)
		return
	}
	var v VitalBeacon
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 32*1024)).Decode(&v); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	v.Name = strings.ToUpper(strings.TrimSpace(v.Name))
	if v.Name == "" || !knownVitals[v.Name] {
		noteRejected()
		http.Error(w, "unknown vital name", http.StatusBadRequest)
		return
	}
	if err := checkVitalValue(v.Name, v.Value); err != nil {
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := checkURL(v.PageURL); err != nil || checkSession(v.SessionID) != nil {
		noteRejected()
		http.Error(w, "invalid vital", http.StatusBadRequest)
		return
	}
	noteAccepted()
	noteVitalClass(classifyVital(v.Name, v.Value))
	if vitals.Add(v) {
		_ = vitals.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}
