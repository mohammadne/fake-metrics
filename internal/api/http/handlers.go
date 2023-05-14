package http

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (handler *Server) liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) simulateOK(c *fiber.Ctx) error {
	_, span := handler.tracer.Start(c.Context(), "api.http.handlers.simulate_ok")
	defer span.End()

	return c.SendStatus(http.StatusOK)
}

func (handler *Server) simulateError(c *fiber.Ctx) error {
	_, span := handler.tracer.Start(c.Context(), "api.http.handlers.simulate_error")
	defer span.End()

	err := errors.New("")
	span.RecordError(err)
	return c.SendStatus(http.StatusForbidden)
}

func (handler *Server) random(c *fiber.Ctx) error {
	_, span := handler.tracer.Start(c.Context(), "api.http.handlers.random")
	defer span.End()

	if rand.Intn(2) == 1 {
		return c.SendStatus(http.StatusOK)
	}

	err := errors.New("")
	span.RecordError(err)
	return c.SendStatus(http.StatusForbidden)
}
