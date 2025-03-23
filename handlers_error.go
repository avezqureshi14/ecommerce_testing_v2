package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var errorEvents = newBatcher(200, 0)
var errorDedupe = newDedupe(2 * time.Second)

func handleErrorBeacon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !wantsJSON(r) {
		noteRejected()
		http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
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
	if errorEvents.Len() >= 5000 {
		noteRejected()
		http.Error(w, "queue full", http.StatusServiceUnavailable)
		return
	}
	if !errorDedupe.Fresh(fingerprint(eb.Message, eb.Stack), time.Now()) {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	noteAccepted()
	eb.Stack = strings.TrimSpace(eb.Stack)
	if len(eb.Stack) > 4000 {
		eb.Stack = eb.Stack[:4000]
	}
	if errorEvents.Add(eb) {
		_ = errorEvents.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}

// trim: callers should TrimSpace message before storing (done in decode next).
