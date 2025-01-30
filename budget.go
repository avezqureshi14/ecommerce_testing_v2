package main

// Budgets follow the public web-vitals cutoffs used by most dashboards.
func classifyVital(name string, value float64) string {
	switch name {
	case "LCP":
		if value <= 2500 {
			return "good"
		}
		if value <= 4000 {
			return "needs-improvement"
		}
		return "poor"
	case "INP":
		if value <= 200 {
			return "good"
		}
		if value <= 500 {
			return "needs-improvement"
		}
		return "poor"
	case "CLS":
		if value <= 0.1 {
			return "good"
		}
		if value <= 0.25 {
			return "needs-improvement"
		}
		return "poor"
	case "FCP":
		if value <= 1800 {
			return "good"
		}
		if value <= 3000 {
			return "needs-improvement"
		}
		return "poor"
	case "TTFB":
		if value <= 800 {
			return "good"
		}
		if value <= 1800 {
			return "needs-improvement"
		}
		return "poor"
	default:
		return "unknown"
	}
}
