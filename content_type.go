package main

import (
	"net/http"
	"strings"
)

func wantsJSON(r *http.Request) bool {
	ct := r.Header.Get("Content-Type")
	if ct == "" {
		return false
	}
	return strings.HasPrefix(strings.ToLower(ct), "application/json")
}
