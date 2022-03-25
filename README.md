# metricstruct


Small library that allows autoregistering metric structs to prometheus Registrer.
This removes the need to register each metric separately in case you don't want to store metrics in global state 


Example

```
type Metrics struct {
	requestsTotal    *prometheus.CounterVec
	cacheHitTotal    *prometheus.CounterVec
	requestsDuration *prometheus.HistogramVec
}

func RegisterMetrics(metricStore prometheus.Registerer) (*Metrics, error) {
	m := Metrics{
		requestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "requests_total",
			Help:      "Total number of requests",
		}, []string{"api_version", "method", "code"}),

		cacheHitTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "cache_hit_total",
			Help:      "Total cache requests",
		}, []string{"type"}),

		requestsDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:      "requests_duration_s",
				Buckets:   prometheus.DefBuckets,
				Help:      "Requst duration in seconds",
			},
			[]string{"api_version"},
		),
	}

	if metricStore == nil {
		return &m, nil
	}

	return &m, metricstruct.Register(metricStore, &m)
}

func newMetricsRegistry() *prometheus.Registry {
	r := prometheus.NewRegistry()
	r.MustRegister(collectors.NewGoCollector())
	r.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	return r
}

func main() {

	metricsRegistry := newMetricsRegistry()
	
	metrics, err := RegisterMetrics(metricsRegistry)

	// check error
	// pass metrics where they are needed
	
}
```

