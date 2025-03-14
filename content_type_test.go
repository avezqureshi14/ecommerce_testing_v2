package main

import (
	"net/http"
	"testing"
)

func TestWantsJSON(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "/", nil)
	if wantsJSON(r) {
		t.Fatal("missing type")
	}
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	if !wantsJSON(r) {
		t.Fatal("json with charset")
	}
	r.Header.Set("Content-Type", "text/plain")
	if wantsJSON(r) {
		t.Fatal("plain text")
	}
}
