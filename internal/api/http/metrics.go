package http

import "github.com/prometheus/client_golang/prometheus"

type Metrics interface {
	IncrementTotalRequests(endpoint string)
}

func newMetrics() Metrics {
	totalRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "total_requests",
		Help: "This is counter for total requests from clients",
	}, []string{"endpoint"})

	// registering metric collectors
	prometheus.MustRegister(totalRequests)

	return &metrics{
		TotalRequests: totalRequests,
	}
}

type metrics struct {
	TotalRequests *prometheus.CounterVec
}

func (m *metrics) IncrementTotalRequests(endpoint string) {
	m.TotalRequests.With(prometheus.Labels{
		"endpoint": endpoint,
	}).Inc()
}
