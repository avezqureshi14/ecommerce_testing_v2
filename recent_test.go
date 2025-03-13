package main

import "testing"

func TestRecentDropsOldest(t *testing.T) {
	r := newRecent(2)
	r.Push("a")
	r.Push("b")
	r.Push("c")
	got := r.Snapshot()
	if len(got) != 2 || got[0] != "b" || got[1] != "c" {
		t.Fatalf("snapshot %#v", got)
	}
}
