package httphandler

import (
	"blockstorage/config"
	"fmt"
	"net/http"
	"sync/atomic"
)

var Concurrency int64 = 0

func Handler(w http.ResponseWriter, r *http.Request) {
	limit := atomic.LoadInt64(&Concurrency)
	if limit < config.Config.MaxConcurrency {
		atomic.AddInt64(&Concurrency, 1)
		switch r.Method {
		case "GET":
			GetHandler(w, r)
		case "POST":
			PostHandler(w, r)
		default:
			fmt.Fprintf(w, "URL.Path = %q , Method = %q \n", r.URL.Path, r.Method)
		}
		atomic.AddInt64(&Concurrency, -1)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unaviable")
	}
}
