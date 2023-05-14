package http

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Server struct {
	config  *Config
	logger  *zap.Logger
	metrics *metrics
	tracer  trace.Tracer
	app     *fiber.App
}

func NewServer(cfg *Config, log *zap.Logger, tracer trace.Tracer) *Server {
	server := &Server{config: cfg, logger: log, metrics: newMetrics(), tracer: tracer}

	server.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.app.Use(server.sharedMetrics)

	server.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	server.app.Get("/healthz/liveness", server.liveness)
	server.app.Get("/healthz/readiness", server.readiness)

	server.app.Get("/ok", server.simulateOK)
	server.app.Get("/slow-ok", server.simulateSlowOK)
	server.app.Get("/error", server.simulateError)
	server.app.Get("/slow-error", server.simulateSlowError)
	server.app.Get("/random", server.random)

	return server
}

func (server *Server) Serve() error {
	addr := fmt.Sprintf(":%d", server.config.Port)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Error("error resolving server", zap.Error(err))
		return err
	}
	return nil
}
