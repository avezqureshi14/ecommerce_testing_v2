package main

import "testing"

func TestRollupGroupsBySession(t *testing.T) {
	views := []PageView{{SessionID: "sess-aaaa", PageURL: "https://example.com/a", LoadMs: 10}}
	errs := []ErrorBeacon{{SessionID: "sess-aaaa", PageURL: "https://example.com/a", Message: "x"}}
	vits := []VitalBeacon{{SessionID: "sess-bbbb", PageURL: "https://example.com/b", Name: "LCP", Value: 5000}}
	got := rollup(views, errs, vits)
	if len(got) != 2 {
		t.Fatalf("sessions %d", len(got))
	}
	var a, b SessionRoll
	for _, row := range got {
		if row.SessionID == "sess-aaaa" {
			a = row
		}
		if row.SessionID == "sess-bbbb" {
			b = row
		}
	}
	if a.Views != 1 || a.Errors != 1 {
		t.Fatalf("a %+v", a)
	}
	if b.Poor != 1 {
		t.Fatalf("b %+v", b)
	}
}
