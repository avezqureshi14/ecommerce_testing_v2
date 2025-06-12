# browser-monitor

Requires Go 1.22.

Small ingest demo for page-view, error, and web-vital beacons.
Nothing here talks to a browser extension or stores cookies.
It keeps a short in-memory batch and a ring of recent pageviews. Restart clears both.

```
go run .
curl -s localhost:8080/healthz
curl -s -X POST localhost:8080/v1/beacon/pageview -H "content-type: application/json" -d "{\"page_url\":\"https://example.com/a\",\"session_id\":\"sess-1234\",\"load_ms\":420}"
```

Batch route: POST /v1/beacon/pageview/batch with {"events":[...]}.

The process idles connections for 75s and shuts down on interrupt.

Env: PORT, MAX_BATCH, MAX_BATCH_AGE_MS.
