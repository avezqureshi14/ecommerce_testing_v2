package main

import (
	"os"
	"strconv"
	"time"
)

// Config keeps the knobs in one place.
type Config struct {
	Addr         string
	MaxBatchSize int
	MaxBatchAge  time.Duration
}

func loadConfig() Config {
	c := Config{Addr: ":8080", MaxBatchSize: 200, MaxBatchAge: 5 * time.Second}
	if v := os.Getenv("PORT"); v != "" {
		c.Addr = ":" + v
	}
	if v := os.Getenv("MAX_BATCH"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			c.MaxBatchSize = n
		}
	}
	return c
}

// env: PORT, MAX_BATCH. MAX_BATCH_AGE stays fixed for now.
