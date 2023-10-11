package repository

import (
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/entity"
	"github.com/NikitaBarysh/metrics_and_alertinc/internal/repository/postgres"
	storage2 "github.com/NikitaBarysh/metrics_and_alertinc/internal/repository/storage"
)

type Storage interface {
	UpdateGaugeMetric(key string, value float64)
	UpdateCounterMetric(key string, value int64)
	GetAllMetric() []entity.Metric
	GetMetric(key string) (entity.Metric, error)
}

func New(DataBaseDSN string) Storage {
	if DataBaseDSN != "" {
		return postgres.NewDBStorage()
	} else {
		return storage2.NewMemStorage()
	}
}
