package main

import (
	"testing"
	"time"
)

func TestBatcherFlushAtCapacity(t *testing.T) {
	b := newBatcher(3, time.Minute)
	if b.Add(1) {
		t.Fatal("should not flush yet")
	}
	if b.Add(2) {
		t.Fatal("should not flush yet")
	}
	if !b.Add(3) {
		t.Fatal("should flush at capacity")
	}
	if got := len(b.Drain()); got != 3 {
		t.Fatalf("got %d events", got)
	}
}

// timer: Drain refreshes lastPush so age-based flush restarts cleanly.
