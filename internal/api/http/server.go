package http

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type Server struct {
	config  *Config
	logger  *zap.Logger
	metrics Metrics
	app     *fiber.App
}

func New(cfg *Config, log *zap.Logger) *Server {
	server := &Server{config: cfg, logger: log}

	server.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	server.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	server.app.Get("/healthz/liveness", server.liveness)
	server.app.Get("/healthz/readiness", server.readiness)

	// v1 := server.app.Group("api/v1")

	// auth := v1.Group("auth")
	// auth.Post("/register", server.register)
	// auth.Post("/login", server.login)

	// contacts := v1.Group("contacts", server.fetchUserId)
	// contacts.Get("/", server.getContacts)
	// contacts.Post("/", server.createContact)
	// contacts.Get("/:id", server.getContact)
	// contacts.Put("/:id", server.updateContact)
	// contacts.Delete("/:id", server.deleteContact)

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
