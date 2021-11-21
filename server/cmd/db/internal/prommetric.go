package internal

import (
	"photographer/internal/prommetric"

	"github.com/prometheus/client_golang/prometheus"
)

// DBPrommetric promtheus mertic for service service
type DBPrommetric struct {
	*prommetric.Prommetric

	// connect to db
	DB_Conn_Fail_Total *prommetric.CounterVec

	// sql crud operations fails, label crud shoud replace with current operation
	DB_CRUD_Fail_Total *prommetric.CounterVec
}

func (service *DB_SERVICE) NewDBPrommetric() {
	baseMetrics := prommetric.NewPrommetric()
	service.Prommetric = &DBPrommetric{
		Prommetric: baseMetrics,

		DB_Conn_Fail_Total: &prommetric.CounterVec{CounterVec: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "db_conn_fail_total",
				Help: "неудачные попытки подключения к базе данных",
			},
			[]string{"date"},
		)},
		DB_CRUD_Fail_Total: &prommetric.CounterVec{CounterVec: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "db_crud_fail_total",
				Help: "неудачные попытки sql операции crud",
			},
			[]string{"crud", "date"},
		)},
	}

	prometheus.MustRegister(service.Prommetric.DB_Conn_Fail_Total)
	prometheus.MustRegister(service.Prommetric.DB_CRUD_Fail_Total)
}
