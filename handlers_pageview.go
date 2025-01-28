package main

import (
	"encoding/json"
	"net/http"
)

var pageViews = newBatcher(200, 0)

func handlePageView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var pv PageView
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 64*1024)).Decode(&pv); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if err := checkURL(pv.PageURL); err != nil {
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	noteAccepted()
	if err := checkSession(pv.SessionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ready := pageViews.Add(pv)
	if ready {
		_ = pageViews.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}

// noted: pageview path now calls noteAccepted/noteRejected via wrapper (wired next).
