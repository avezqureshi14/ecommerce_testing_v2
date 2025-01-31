package main

import "testing"

func TestClassifyVitalBoundaries(t *testing.T) {
	if classifyVital("LCP", 2500) != "good" {
		t.Fatal("lcp 2500 should be good")
	}
	if classifyVital("LCP", 2501) != "needs-improvement" {
		t.Fatal("lcp just over 2500")
	}
	if classifyVital("CLS", 0.25) != "needs-improvement" {
		t.Fatal("cls 0.25 is the poor threshold, not past it")
	}
	if classifyVital("CLS", 0.26) != "poor" {
		t.Fatal("cls 0.26")
	}
	if classifyVital("nope", 1) != "unknown" {
		t.Fatal("unknown name")
	}
}
