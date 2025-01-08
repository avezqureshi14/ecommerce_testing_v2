package main

import (
	"sync"
	"time"
)

// Batcher holds decoded events until flush.
type Batcher struct {
	mu       sync.Mutex
	events   []any
	maxItems int
	maxAge   time.Duration
	lastPush time.Time
}

func newBatcher(maxItems int, maxAge time.Duration) *Batcher {
	return &Batcher{maxItems: maxItems, maxAge: maxAge, lastPush: time.Now()}
}

// Add returns true when the batch is ready to flush.
func (b *Batcher) Add(ev any) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.events = append(b.events, ev)
	if len(b.events) >= b.maxItems {
		return true
	}
	return time.Since(b.lastPush) >= b.maxAge
}

func (b *Batcher) Drain() []any {
	b.mu.Lock()
	defer b.mu.Unlock()
	out := b.events
	b.events = nil
	b.lastPush = time.Now()
	return out
}

func (b *Batcher) Len() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.events)
}
