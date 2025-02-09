package main

import "net/http"

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/v1/recent", handleRecent)
	mux.HandleFunc("/v1/beacon/pageview", handlePageView)
	mux.HandleFunc("/v1/beacon/error", handleErrorBeacon)
	mux.HandleFunc("/v1/beacon/vitals", handleVitals)
	return mux
}

// routes: /health, /v1/beacon/pageview, /v1/beacon/error, /v1/beacon/vitals
