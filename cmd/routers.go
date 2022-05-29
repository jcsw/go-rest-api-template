package main

import (
	hello "go-rest-api-template/pkg/hello"
	monitor "go-rest-api-template/pkg/monitor"
	sys "go-rest-api-template/pkg/system"
	http "net/http"
	atomic "sync/atomic"
)

func handleRouters(mux *http.ServeMux) {
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/monitor", monitor.Handler)
	mux.HandleFunc("/hello", hello.Handler)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sys.HTTPResponseWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	if atomic.LoadInt32(&healthy) == 1 {
		sys.HTTPResponseWithCode(w, http.StatusOK)
		return
	}

	sys.HTTPResponseWithCode(w, http.StatusServiceUnavailable)
}
