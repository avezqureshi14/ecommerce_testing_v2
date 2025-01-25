package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	fmt.Println("browser-monitor starting on :8080")
	http.ListenAndServe(":8080", nil)
}


// startup banner for ops visibility


// boot: loadConfig drives Addr and batch limits (wired in Feb batch).
