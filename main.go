package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Counter metric to track total HTTP requests
var requests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "app_requests_total",
		Help: "Total number of HTTP requests received",
	},
)

// Root handler increments counter and returns response
func handler(w http.ResponseWriter, r *http.Request) {
	requests.Inc()
	fmt.Fprintf(w, "Student Platform Running")
}

func main() {

	// Register Prometheus metrics
	prometheus.MustRegister(requests)

	// Define routes
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	// Start HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}