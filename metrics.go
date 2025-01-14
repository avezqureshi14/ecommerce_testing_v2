package main

import (
	"sync/atomic"
)

// Counters are read by /health and tests.
var (
	acceptedCount atomic.Int64
	rejectedCount atomic.Int64
)

func noteAccepted() { acceptedCount.Add(1) }
func noteRejected() { rejectedCount.Add(1) }
