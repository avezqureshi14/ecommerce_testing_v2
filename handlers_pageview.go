package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var pageViews = newBatcher(200, 0)

func handlePageView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !wantsJSON(r) {
		noteRejected()
		http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	if !sessionLimit.Allow(r.Header.Get("X-Session")+r.RemoteAddr, time.Now()) {
		noteLimited()
		http.Error(w, "slow down", http.StatusTooManyRequests)
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
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := checkLoad(pv.LoadMs); err != nil {
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := checkNav(pv.NavType); err != nil {
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := checkReferrer(pv.Referrer); err != nil {
		noteRejected()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pv.UserAgent = clipUA(pv.UserAgent)
	if pageViews.Len() >= maxQueued {
		noteRejected()
		http.Error(w, "queue full", http.StatusServiceUnavailable)
		return
	}
	recentPageViews.Push(pv)
	ready := pageViews.Add(pv)
	if ready {
		_ = pageViews.Drain()
	}
	w.WriteHeader(http.StatusAccepted)
}

// noted: pageview path now calls noteAccepted/noteRejected via wrapper (wired next).
