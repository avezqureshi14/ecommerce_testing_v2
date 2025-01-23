package main

import "testing"

func TestCounters(t *testing.T) {
	before := acceptedCount.Load()
	noteAccepted()
	if acceptedCount.Load() != before+1 {
		t.Fatal("accepted counter did not move")
	}
	noteRejected()
	if rejectedCount.Load() < 1 {
		t.Fatal("rejected counter did not move")
	}
}
