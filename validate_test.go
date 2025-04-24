package main

import (
	"strings"
	"testing"
)

var strings2 = strings

func TestCheckURL(t *testing.T) {
	if err := checkURL(""); err == nil {
		t.Fatal("expected error for empty url")
	}
	if err := checkURL("notaurl"); err == nil {
		t.Fatal("expected error for bad url")
	}
	if err := checkURL("https://example.com/a"); err != nil {
		t.Fatal(err)
	}
}

func TestCheckSession(t *testing.T) {
	if err := checkSession("short"); err == nil {
		t.Fatal("expected error for short session")
	}
	if err := checkSession("sess-12345678"); err != nil {
		t.Fatal(err)
	}
}

// http scheme is accepted alongside https.

func TestClipUA(t *testing.T) {
	if clipUA("  hi  ") != "hi" {
		t.Fatal("trim")
	}
	long := strings2.Repeat("a", 200)
	if len(clipUA(long)) != 180 {
		t.Fatal("clip")
	}
}

func TestCheckLoad(t *testing.T) {
	if err := checkLoad(0); err != nil {
		t.Fatal(err)
	}
	if checkLoad(-1) == nil {
		t.Fatal("negative load")
	}
	if checkLoad(6*60*1000) == nil {
		t.Fatal("runaway load")
	}
}

func TestCheckNavAndReferrer(t *testing.T) {
	if err := checkNav(""); err != nil {
		t.Fatal(err)
	}
	if err := checkNav("reload"); err != nil {
		t.Fatal(err)
	}
	if checkNav("typed") == nil {
		t.Fatal("typed is not stored")
	}
	if err := checkReferrer(""); err != nil {
		t.Fatal(err)
	}
	if checkReferrer("notaurl") == nil {
		t.Fatal("bad referrer")
	}
}
