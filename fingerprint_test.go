package main

import "testing"

func TestFingerprintIgnoresLineNumbers(t *testing.T) {
	a := fingerprint("boom at line 12", "at render (app.js:12:3)")
	b := fingerprint("boom at line 40", "at render (app.js:40:3)")
	if a != b {
		t.Fatalf("wanted same group, got %s vs %s", a, b)
	}
	c := fingerprint("other failure", "at render (app.js:12:3)")
	if a == c {
		t.Fatal("different messages should not collapse")
	}
	if fingerprint("only message", "") == "" {
		t.Fatal("empty stack still has a hash")
	}
}
