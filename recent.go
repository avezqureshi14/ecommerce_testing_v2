package main

import "sync"

// Recent is a fixed ring of accepted payloads.
type Recent struct {
	mu    sync.Mutex
	buf   []any
	next  int
	count int
}

func newRecent(n int) *Recent {
	if n < 1 {
		n = 1
	}
	return &Recent{buf: make([]any, n)}
}

func (r *Recent) Push(ev any) {
	r.mu.Lock()
	r.buf[r.next] = ev
	r.next = (r.next + 1) % len(r.buf)
	if r.count < len(r.buf) {
		r.count++
	}
	r.mu.Unlock()
}

func (r *Recent) Snapshot() []any {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]any, 0, r.count)
	start := 0
	if r.count == len(r.buf) {
		start = r.next
	}
	for i := 0; i < r.count; i++ {
		idx := (start + i) % len(r.buf)
		out = append(out, r.buf[idx])
	}
	return out
}
