package config

import (
	"github.com/MohammadNE/fake-metrics/internal/api/http"
	"github.com/MohammadNE/fake-metrics/pkg/logger"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		HTTP: &http.Config{
			Address: "localhost",
			Port:    8080,
		},
	}
}
