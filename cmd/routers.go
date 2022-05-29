package main

import (
	hello "go-rest-api-template/pkg/hello"
	monitor "go-rest-api-template/pkg/monitor"
	http "net/http"
	atomic "sync/atomic"
)

func handleRouters(mux *http.ServeMux) {
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/monitor", monitor.Handler)
	mux.HandleFunc("/hello", hello.Handler)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}
