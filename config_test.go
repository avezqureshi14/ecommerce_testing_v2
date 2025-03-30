package main

import (
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("MAX_BATCH", "")
	c := loadConfig()
	if c.Addr != ":8080" || c.MaxBatchSize != 200 {
		t.Fatalf("bad defaults: %+v", c)
	}
}

func TestLoadConfigBadBatchIgnored(t *testing.T) {
	t.Setenv("MAX_BATCH", "nope")
	c := loadConfig()
	if c.MaxBatchSize != 200 {
		t.Fatalf("should ignore bad value: %+v", c)
	}
	t.Setenv("MAX_BATCH_AGE_MS", "2500")
	if loadConfig().MaxBatchAge.Milliseconds() != 2500 {
		t.Fatal("age")
	}
}
