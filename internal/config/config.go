package config

import (
	"github.com/mohammadne/fake-metrics/internal/api/http"
	"github.com/mohammadne/fake-metrics/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	HTTP   *http.Config   `koanf:"http"`
}
