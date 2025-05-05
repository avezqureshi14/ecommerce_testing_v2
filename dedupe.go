package main

import (
	"sync"
	"time"
)

type seenKey struct {
	fp  string
	exp time.Time
}

// Dedupe drops a repeat of the same fingerprint inside the window.
type Dedupe struct {
	mu      sync.Mutex
	window  time.Duration
	seen    map[string]time.Time
}

func newDedupe(window time.Duration) *Dedupe {
	return &Dedupe{window: window, seen: map[string]time.Time{}}
}

func (d *Dedupe) Fresh(fp string, now time.Time) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if exp, ok := d.seen[fp]; ok && now.Before(exp) {
		return false
	}
	d.seen[fp] = now.Add(d.window)
	if len(d.seen) > 4000 {
		for k, exp := range d.seen {
			if !now.Before(exp) {
				delete(d.seen, k)
			}
		}
	}
	return true
}
