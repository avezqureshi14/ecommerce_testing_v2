package main

// SessionRoll keeps a tiny per-session tally for the demo UI.
type SessionRoll struct {
	SessionID string
	Views     int
	Errors    int
	Vitals    int
	Poor      int
}

func rollup(views []PageView, errors []ErrorBeacon, vitals []VitalBeacon) []SessionRoll {
	by := map[string]*SessionRoll{}
	touch := func(id string) *SessionRoll {
		r, ok := by[id]
		if !ok {
			r = &SessionRoll{SessionID: id}
			by[id] = r
		}
		return r
	}
	for _, v := range views {
		touch(v.SessionID).Views++
	}
	for _, e := range errors {
		touch(e.SessionID).Errors++
	}
	for _, v := range vitals {
		row := touch(v.SessionID)
		row.Vitals++
		if v.Rating == "poor" || classifyVital(v.Name, v.Value) == "poor" {
			row.Poor++
		}
	}
	out := make([]SessionRoll, 0, len(by))
	for _, r := range by {
		out = append(out, *r)
	}
	return out
}
