package main

// PageView is sent once per navigation.
type PageView struct {
	PageURL   string `json:"page_url"`
	SessionID string `json:"session_id"`
	Referrer  string `json:"referrer,omitempty"`
	LoadMs    int64  `json:"load_ms"`
	UserAgent string `json:"user_agent,omitempty"`
	NavType   string `json:"nav_type,omitempty"`
}

// ErrorBeacon captures a client-side failure.
type ErrorBeacon struct {
	PageURL   string `json:"page_url"`
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
	Stack     string `json:"stack,omitempty"`
}

// VitalBeacon carries one web-vital sample.
type VitalBeacon struct {
	PageURL   string  `json:"page_url"`
	SessionID string  `json:"session_id"`
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
}

// SessionID is issued by the demo snippet and must be at least 8 chars.
