package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	cfg := loadConfig()
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      newMux(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Println("browser-monitor listening on", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server error:", err)
	}
}
