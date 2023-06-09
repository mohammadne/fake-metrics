package cmd

import (
	"os"

	"github.com/mohammadne/fake-metrics/internal/api/http"
	"github.com/mohammadne/fake-metrics/internal/config"
	"github.com/mohammadne/fake-metrics/pkg/logger"
	"github.com/mohammadne/fake-metrics/pkg/tracing"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run HTTP server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	logger := logger.New(cfg.Logger)
	tracer := tracing.New(cfg.Tracing)

	server := http.NewServer(cfg.HTTP, logger, tracer)
	go server.Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by receiving a unix signal", field)
}
