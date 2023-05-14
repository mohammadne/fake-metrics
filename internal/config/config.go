package config

import (
	"github.com/mohammadne/fake-metrics/internal/api/http"
	"github.com/mohammadne/fake-metrics/pkg/logger"
	"github.com/mohammadne/fake-metrics/pkg/tracing"
)

type Config struct {
	Logger  *logger.Config `koanf:"logger"`
	Tracing tracing.Config `koanf:"tracing"`
	HTTP    *http.Config   `koanf:"http"`
}
