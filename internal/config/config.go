package config

import (
	"github.com/MohammadNE/fake-metrics/internal/api/http"
	"github.com/MohammadNE/fake-metrics/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	HTTP   *http.Config   `koanf:"http"`
}
