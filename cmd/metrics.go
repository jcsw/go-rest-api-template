package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestDurationMetric = createHttpRequestDurationMetric()
	httpRequestCountMetric    = createHttpRequestCountMetric()
)

func initMetrics() {
	prometheus.Register(httpRequestDurationMetric)
	prometheus.Register(httpRequestCountMetric)
}

func createHttpRequestDurationMetric() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests",
		Buckets:   prometheus.DefBuckets,
	}, []string{"uri", "method", "code"})
}

func createHttpRequestCountMetric() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http",
		Name:      "request_count",
		Help:      "The count of the HTTP requests",
	}, []string{"uri", "method", "code"})
}

func httpRequestMetric(uri string, method string, status string, duration float64) {
	httpRequestDurationMetric.WithLabelValues(uri, method, status).Observe(duration)
	httpRequestCountMetric.WithLabelValues(uri, method, status).Add(1)
}
