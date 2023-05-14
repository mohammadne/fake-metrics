package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *Server) liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) ok(c *fiber.Ctx) error {
	handler.metrics.IncrementTotalRequests("ok")
	return c.SendStatus(http.StatusOK)
}
