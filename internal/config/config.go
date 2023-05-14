package config

import (
	"github.com/MohammadNE/fake-metrics/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
