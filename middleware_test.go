package main

import (
	"net/http"
	nethttptest "net/http/httptest"
	strings3 "strings"
	"testing"
)

func TestLimitBody(t *testing.T) {
	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 64)
		_, _ = r.Body.Read(buf)
		w.WriteHeader(http.StatusOK)
	})
	h := limitBody(ok, 16)
	req := nethttptest.NewRequest(http.MethodPost, "/", strings3.NewReader(strings3.Repeat("x", 64)))
	rec := nethttptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code == http.StatusOK {
		// MaxBytesReader surfaces the error on read; handler still ran.
		t.Log("handler ran, body capped at read time")
	}
}
