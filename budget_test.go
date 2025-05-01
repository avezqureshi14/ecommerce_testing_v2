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
	if classifyVital("FCP", 1801) != "needs-improvement" {
		t.Fatal("fcp boundary")
	}
	if classifyVital("TTFB", 800) != "good" {
		t.Fatal("ttfb good edge")
	}
	if classifyVital("LCP", 0) != "good" {
		t.Fatal("lcp zero")
	}
	if classifyVital("CLS", 0) != "good" {
		t.Fatal("cls zero")
	}
	if classifyVital("INP", 500) != "needs-improvement" {
		t.Fatal("inp 500")
	}
	if classifyVital("INP", 501) != "poor" {
		t.Fatal("inp 501")
	}
	if classifyVital("nope", 1) != "unknown" {
		t.Fatal("unknown name")
	}
}
