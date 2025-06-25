package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalRequests =  prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_req_TOTAL",
			Help: "Всего HTTP запросов",
		},
		[]string{"path", "method"},
	)

	HTTPDuration =  prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_duration_SECOND",
			Help: "Время обработки ENDPOINTS",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	CacheDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cache_duration_SECOND",
			Help: "Время работы с Редис",
		}, []string{"method"},
	)
	APIDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "API_duration_SECOND",
			Help: "Время работы с внешним API",
		}, []string{"method"},
	)
)