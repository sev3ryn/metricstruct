package metricstruct

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestRegister(t *testing.T) {
	r := prometheus.NewRegistry()

	metrics := struct {
		counter    prometheus.Counter
		counterVec *prometheus.CounterVec
	}{
		counter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "counter",
			Help: "counter help",
		}),
		counterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "counter_vec",
			Help: "counter vec help",
		}, []string{"label"}),
	}

	if err := Register(r, &metrics); err != nil {
		t.Fatalf("Register() = %v", err)
	}

	if !r.Unregister(metrics.counter) {
		t.Fatalf("Unregister() = counter doesn't exist")
	}
	if !r.Unregister(metrics.counterVec) {
		t.Fatalf("Unregister() = counterVec doesn't exist")
	}
}
