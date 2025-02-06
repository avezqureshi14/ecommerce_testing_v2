package main

import (
	"sync"
	"time"
)

type bucket struct {
	tokens float64
	last   time.Time
}

// Limiter is a tiny per-session token bucket. Not distributed.
type Limiter struct {
	mu     sync.Mutex
	rate   float64
	burst  float64
	byKey  map[string]*bucket
}

func newLimiter(perSec float64, burst int) *Limiter {
	return &Limiter{rate: perSec, burst: float64(burst), byKey: map[string]*bucket{}}
}

func (l *Limiter) Allow(key string, now time.Time) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	b, ok := l.byKey[key]
	if !ok {
		b = &bucket{tokens: l.burst, last: now}
		l.byKey[key] = b
	}
	elapsed := now.Sub(b.last).Seconds()
	b.tokens += elapsed * l.rate
	if b.tokens > l.burst {
		b.tokens = l.burst
	}
	b.last = now
	if b.tokens < 1 {
		return false
	}
	b.tokens -= 1
	return true
}

var sessionLimit = newLimiter(20, 40)
