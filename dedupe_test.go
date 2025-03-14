package main

import (
	"testing"
	"time"
)

func TestDedupeWindow(t *testing.T) {
	d := newDedupe(time.Second)
	now := time.Date(2025, 3, 4, 9, 12, 0, 0, time.UTC)
	if !d.Fresh("abc", now) {
		t.Fatal("first")
	}
	if d.Fresh("abc", now.Add(200*time.Millisecond)) {
		t.Fatal("still inside the window")
	}
	if !d.Fresh("abc", now.Add(2*time.Second)) {
		t.Fatal("window passed")
	}
}
