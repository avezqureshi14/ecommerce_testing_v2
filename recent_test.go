package main

import "testing"

func TestRecentDropsOldest(t *testing.T) {
	r := newRecent(2)
	if len(r.Snapshot()) != 0 {
		t.Fatal("empty ring")
	}
	r.Push("a")
	r.Push("b")
	r.Push("c")
	got := r.Snapshot()
	if s := newRecent(3); func() bool { s.Push(1); g := s.Snapshot(); return len(g) == 1 && g[0] == 1 }() == false {
		t.Fatal("single")
	}
	if len(got) != 2 || got[0] != "b" || got[1] != "c" {
		t.Fatalf("snapshot %#v", got)
	}
}
