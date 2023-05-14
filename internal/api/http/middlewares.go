package http

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (handler *Server) sharedMetrics(c *fiber.Ctx) error {
	start := time.Now()
	method := c.Route().Method
	path := c.Route().Path

	if path == "/metrics" {
		return c.Next()
	}

	handler.metrics.RequestsInProgress.WithLabelValues(method, path).Inc()
	defer func() {
		handler.metrics.RequestsInProgress.WithLabelValues(method, path).Dec()
	}()

	status := fiber.StatusInternalServerError
	err := c.Next() // process the actual handler
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			// Get correct error code from fiber.Error type
			status = e.Code
		}
	} else {
		status = c.Response().StatusCode()
	}

	statusCode := strconv.Itoa(status)
	handler.metrics.RequestsInProgress.WithLabelValues(statusCode, method, path).Inc()

	elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
	handler.metrics.RequestsDuration.WithLabelValues(statusCode, method, path).Observe(elapsed)

	return err
}
