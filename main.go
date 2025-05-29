package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := loadConfig()
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      withCORS(newMux()),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  75 * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go runFlush(ctx, pageViews, cfg.MaxBatchAge)
	go runFlush(ctx, errorEvents, cfg.MaxBatchAge)
	go runFlush(ctx, vitals, cfg.MaxBatchAge)
	fmt.Println("browser-monitor listening on", cfg.Addr)
	errCh := make(chan error, 1)
	go func() { errCh <- srv.ListenAndServe() }()
	select {
	case <-ctx.Done():
		// five seconds is enough to drain the flush loops on a laptop
		shut, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shut)
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}
}
