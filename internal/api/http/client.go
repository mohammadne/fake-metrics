package http

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type client struct {
	config *Config
	logger *zap.Logger
	app    *fiber.Client
}

func NewClient(cfg *Config, log *zap.Logger) *client {
	return &client{config: cfg, logger: log, app: fiber.AcquireClient()}
}

func (c *client) Request(wg *sync.WaitGroup, path string) {
	defer wg.Done()

	uri := fmt.Sprintf("%s:%d/%s", c.config.Address, c.config.Port, path)
	agent := c.app.Get(uri)

	if err := agent.Parse(); err != nil {
		panic(err)
	}

	code, body, errs := agent.Bytes()
	fmt.Println(code, body, errs)
}
