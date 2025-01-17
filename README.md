# browser-monitor

Demo telemetry ingest: page-view, error, and web-vital beacons over HTTP.

```sh
go test ./...
go run . 
curl localhost:8080/health
```

`POST /v1/beacon/pageview`, `/v1/beacon/error`, `/v1/beacon/vitals` accept JSON
up to 64KB. Events batch in memory; `/health` reports depth.
