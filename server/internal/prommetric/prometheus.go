package prommetric

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// CounterVec wrapper for further using
type CounterVec struct {
	*prometheus.CounterVec
}

// GaugeVec wrapper for further using
type GaugeVec struct {
	*prometheus.GaugeVec
}

// Prommetric general prometheus metric
type Prommetric struct {
	// consul access fails
	CONSUL_ACCESS_FAIL_TOTAL *CounterVec

	// service access fails
	SERVICE_ACCESS_FAIL_TOTAL *CounterVec
}

// Add date label with current date to metric
func (v *CounterVec) WithLabelValues(lvs ...string) prometheus.Counter {
	lvs = append(lvs, strconv.Itoa(int(time.Now().Unix())))
	return v.CounterVec.WithLabelValues(lvs...)
}

// PromCounterVec returns embedded CounterVec
func (v *CounterVec) PromCounterVec() *prometheus.CounterVec {
	return v.CounterVec
}

// Add date label with current date to metric
func (v *GaugeVec) WithLabelValues(lvs ...string) prometheus.Counter {
	lvs = append(lvs, strconv.Itoa(int(time.Now().Unix())))
	return v.GaugeVec.WithLabelValues(lvs...)
}

// PromGaugeVec returns embedded GaugeVec
func (v *GaugeVec) PromGaugeVec() *prometheus.GaugeVec {
	return v.GaugeVec
}

// NewPrommetric return new prometheus metric
func NewPrommetric() *Prommetric {
	metrics := &Prommetric{
		CONSUL_ACCESS_FAIL_TOTAL: &CounterVec{prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "consul_access_fail_total",
				Help: "неудачные попытки доступа к консулу",
			},
			[]string{"date"},
		)},
		SERVICE_ACCESS_FAIL_TOTAL: &CounterVec{prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "service_access_fail_total",
				Help: "неудачные попытки доступа к сервису service_name",
			},
			[]string{"service_name", "date"},
		)},
	}

	// register to prometheus
	prometheus.MustRegister(metrics.CONSUL_ACCESS_FAIL_TOTAL.MetricVec)
	prometheus.MustRegister(metrics.SERVICE_ACCESS_FAIL_TOTAL)
	return metrics
}
