package main

import "testing"

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
