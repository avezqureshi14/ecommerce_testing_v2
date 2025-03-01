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

var (
	vitalGood    atomic.Int64
	vitalPoor    atomic.Int64
)

func noteVitalClass(class string) {
	switch class {
	case "good":
		vitalGood.Add(1)
	case "poor":
		vitalPoor.Add(1)
	}
}

var limitedCount atomic.Int64

func noteLimited() { limitedCount.Add(1) }
