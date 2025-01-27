package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Define a counter metric
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests processed by the MyApp web server.",
		},
		[]string{"method", "path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_errors_total",
			Help: "Total number of error requests processed by the MyApp web server.",
		},
		[]string{"method", "path", "status"},
	)

	// Define a histogram metric
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for handler in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// Request size in bytes
	RequestSizeBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "Size of HTTP requests in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 6), // 100B to ~1MB
		},
		[]string{"method", "path"},
	)

	// Response size in bytes
	ResponseSizeBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 6), // 100B to ~1MB
		},
		[]string{"method", "path"},
	)
)

// handle save global configuration
func saveToGlobal() {
	global.Prometheus = &setting.PrometheusSetting{}

	global.Prometheus.RequestCount = RequestCount
	global.Prometheus.ErrorCount = ErrorCount
	global.Prometheus.RequestDuration = RequestDuration
	global.Prometheus.RequestSizeBytes = RequestSizeBytes
	global.Prometheus.ResponseSizeBytes = ResponseSizeBytes
}

// handle register metric collection
func RegisterMetrics() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestSizeBytes)
	prometheus.MustRegister(ResponseSizeBytes)
}

// handle initialize global variables
func InitPrometheus() {
	// register metric collection
	RegisterMetrics()
	// save global variables
	saveToGlobal()
}
