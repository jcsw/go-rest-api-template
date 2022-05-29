package main

import (
	hello "go-rest-api-template/pkg/hello"
	monitor "go-rest-api-template/pkg/infra/monitor"
	sys "go-rest-api-template/pkg/infra/system"

	http "net/http"
	regexp "regexp"
	strings "strings"
	atomic "sync/atomic"

	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute("GET", "/ping", ping),
	newRoute("GET", "/monitor", monitor.Get),
	newRoute("GET", "/hello", hello.Get),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func handleRouters(mux *http.ServeMux) {
	mux.HandleFunc("/", buildRouters)
	mux.Handle("/metrics", promhttp.Handler())
}

func buildRouters(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			route.handler(w, r)
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ","))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		sys.HTTPResponseWithCode(w, http.StatusOK)
		return
	}

	sys.HTTPResponseWithCode(w, http.StatusServiceUnavailable)
}
