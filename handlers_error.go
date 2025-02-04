package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
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
	if err := checkURL(eb.PageURL); err != nil || checkSession(eb.SessionID) != nil {
		noteRejected()
		http.Error(w, "invalid error beacon", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(eb.Message) == "" {
		noteRejected()
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}
	eb.Message = strings.TrimSpace(eb.Message)
	noteAccepted()
	if len(eb.Stack) > 4000 {
		eb.Stack = eb.Stack[:4000]
	}
	if errorEvents.Add(eb) {
		_ = errorEvents.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}

// trim: callers should TrimSpace message before storing (done in decode next).
