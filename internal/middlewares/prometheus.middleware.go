package middlewares

import (
	"net/http"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Capture request size
		requestSize := computeRequestSize(c.Request)
		
		// Process request
		c.Next()
		
		// Calculate metrics
		duration := time.Since(startTime).Seconds()
		status := c.Writer.Status()
		path := c.FullPath() // Get the registered route path, e.g., "/users/:id"
		responseSize := float64(c.Writer.Size())

		global.Prometheus.RequestCount.WithLabelValues(c.Request.Method ,path, http.StatusText(status)).Inc()
		global.Prometheus.RequestDuration.WithLabelValues(c.Request.Method, path).Observe(duration)
		global.Prometheus.RequestSizeBytes.WithLabelValues(c.Request.Method, path).Observe(requestSize)
		global.Prometheus.ResponseSizeBytes.WithLabelValues(c.Request.Method, path).Observe(responseSize)

		// Capture error count
		if status >= 400 {
			global.Prometheus.ErrorCount.WithLabelValues(c.Request.Method, path, http.StatusText(status)).Inc()
		}
	}
}

// Compute request size
func computeRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size += len(r.URL.String())
	}
	if r.Method != "" {
		size += len(r.Method)
	}
	if r.Proto != "" {
		size += len(r.Proto)
	}
	for key, values := range r.Header {
		size += len(key)
		for _, value := range values {
			size += len(value)
		}
	}
	if r.ContentLength > 0 {
		size += int(r.ContentLength)
	}
	return float64(size)
}
