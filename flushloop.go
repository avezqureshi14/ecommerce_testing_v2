package main

import (
	"context"
	"time"
)

func runFlush(ctx context.Context, b *Batcher, every time.Duration) {
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			_ = b.Drain()
			return
		case <-t.C:
			if b.Len() > 0 {
				_ = b.Drain()
			}
		}
	}
}
