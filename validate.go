package main

import (
	"errors"
	"net/url"
	strings2 "strings"
)

func checkURL(raw string) error {
	if strings2.TrimSpace(raw) == "" {
		return errors.New("page_url is required")
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return errors.New("page_url must be http(s)")
	}
	return nil
}

func checkSession(s string) error {
	if len(s) < 8 {
		return errors.New("session_id too short")
	}
	return nil
}

// note: checkURL already permits query strings via ParseRequestURI; no change needed.
