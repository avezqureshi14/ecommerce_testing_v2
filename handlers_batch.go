package main

import (
	"encoding/json"
	"net/http"
)

const maxBatchEvents = 100

type pageBatch struct {
	Events []PageView `json:"events"`
}

func handlePageBatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !wantsJSON(r) {
		noteRejected()
		http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	var body pageBatch
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 256*1024)).Decode(&body); err != nil {
		noteRejected()
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if len(body.Events) == 0 || len(body.Events) > maxBatchEvents {
		noteRejected()
		http.Error(w, "events must be 1..100", http.StatusBadRequest)
		return
	}
	accepted := 0
	for _, pv := range body.Events {
		if checkURL(pv.PageURL) != nil || checkSession(pv.SessionID) != nil || checkLoad(pv.LoadMs) != nil {
			noteRejected()
			continue
		}
		pv.UserAgent = clipUA(pv.UserAgent)
		recentPageViews.Push(pv)
		pageViews.Add(pv)
		noteAccepted()
		accepted++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(map[string]int{"accepted": accepted})
}
