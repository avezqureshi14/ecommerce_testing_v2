# browser-monitor

Small ingest demo for page-view, error, and web-vital beacons.
Nothing here talks to a browser extension or stores cookies.

```
go run .
curl -s localhost:8080/healthz
curl -s -X POST localhost:8080/v1/beacon/pageview -H "content-type: application/json" -d "{\"page_url\":\"https://example.com/a\",\"session_id\":\"sess-1234\",\"load_ms\":420}"
```

Batch route: POST /v1/beacon/pageview/batch with {"events":[...]}.

Env: PORT, MAX_BATCH, MAX_BATCH_AGE_MS.
