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

func checkLoad(ms int64) error {
	if ms < 0 {
		return errors.New("load_ms cannot be negative")
	}
	if ms > 5*60*1000 {
		return errors.New("load_ms looks like it was left running")
	}
	return nil
}

func checkVitalValue(name string, value float64) error {
	if value < 0 {
		return errors.New("vital value cannot be negative")
	}
	if name == "CLS" && value > 10 {
		return errors.New("cls value is not plausible")
	}
	return nil
}

func clipUA(ua string) string {
	ua = strings2.TrimSpace(ua)
	if len(ua) > 180 {
		return ua[:180]
	}
	return ua
}

func checkNav(nav string) error {
	if nav == "" {
		return nil
	}
	switch nav {
	case "navigate", "reload", "back_forward", "prerender":
		return nil
	default:
		return errors.New("nav_type is not one we store")
	}
}
