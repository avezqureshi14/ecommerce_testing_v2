package main

import (
	"encoding/json"
	"net/http"
)

var errorEvents = newBatcher(200, 0)

func handleErrorBeacon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var eb ErrorBeacon
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 64*1024)).Decode(&eb); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if err := checkURL(eb.PageURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(eb.Stack) > 4000 {
		eb.Stack = eb.Stack[:4000]
	}
	if errorEvents.Add(eb) {
		_ = errorEvents.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}

// trim: callers should TrimSpace message before storing (done in decode next).
