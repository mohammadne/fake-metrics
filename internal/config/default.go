package config

import (
	"github.com/mohammadne/fake-metrics/internal/api/http"
	"github.com/mohammadne/fake-metrics/pkg/logger"
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
