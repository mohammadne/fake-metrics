package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalRequests      *prometheus.CounterVec
	RequestsInProgress *prometheus.GaugeVec
	RequestsDuration   *prometheus.HistogramVec
}

func newMetrics() *metrics {
	constLabels := prometheus.Labels{
		"service": "fake-metrics",
	}

	totalRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "total_requests",
		Help:        "This is counter for total requests from clients.",
		ConstLabels: constLabels,
	}, []string{"status_code", "method", "path"})

	requestsInProgress := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "requests_in_progress",
		Help:        "All the requests in progress based on method and path.",
		ConstLabels: constLabels,
	}, []string{"method", "path"})

	requestsDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "request_duration_seconds",
		Help:        "Duration of all HTTP requests by status code, method and path.",
		ConstLabels: constLabels,
		Buckets: []float64{
			0.000000001, // 1ns
			0.000000002,
			0.000000005,
			0.00000001, // 10ns
			0.00000002,
			0.00000005,
			0.0000001, // 100ns
			0.0000002,
			0.0000005,
			0.000001, // 1µs
			0.000002,
			0.000005,
			0.00001, // 10µs
			0.00002,
			0.00005,
			0.0001, // 100µs
			0.0002,
			0.0005,
			0.001, // 1ms
			0.002,
			0.005,
			0.01, // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
			10.0, // 10s
			15.0,
			20.0,
			30.0,
		},
	},
		[]string{"status_code", "method", "path"},
	)

	// registering metric collectors
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(requestsInProgress)
	prometheus.MustRegister(requestsDuration)

	return &metrics{
		TotalRequests:      totalRequests,
		RequestsInProgress: requestsInProgress,
		RequestsDuration:   requestsDuration,
	}
}
