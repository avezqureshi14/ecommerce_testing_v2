package main

import (
	"testing"
	"time"
)

func TestLimiterRefills(t *testing.T) {
	l := newLimiter(10, 1)
	now := time.Date(2025, 3, 2, 11, 4, 0, 0, time.UTC)
	if !l.Allow("s", now) {
		t.Fatal("first token")
	}
	if l.Allow("s", now) {
		t.Fatal("burst is 1")
	}
	later := now.Add(200 * time.Millisecond)
	if !l.Allow("s", later) {
		t.Fatal("should have refilled")
	}
	l2 := newLimiter(1, 1)
	if !l2.Allow("other", now) {
		t.Fatal("other session")
	}
}
