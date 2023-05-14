package cmd

import (
	"sync"

	"github.com/mohammadne/fake-metrics/internal/api/http"
	"github.com/mohammadne/fake-metrics/internal/config"
	"github.com/mohammadne/fake-metrics/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Client struct{}

func (client Client) Command() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) {
		client.main(config.Load(false), cmd, args)
	}

	cmd := &cobra.Command{
		Use:   "client",
		Short: "simulate HTTP requests to the server",
		Run:   run,
		Args:  cobra.OnlyValidArgs,
	}

	cmd.ValidArgs = []string{"ok", "slow-ok", "error", "slow-error", "random"}
	cmd.Flags().Int("count", 100, "Specify how many times process the request.")

	return cmd
}

func (*Client) main(cfg *config.Config, cmd *cobra.Command, args []string) {
	logger := logger.New(cfg.Logger)

	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		logger.Fatal("error retrieving count", zap.Error(err))
	}

	if len(args) != 1 {
		logger.Fatal("invalid arguments given", zap.Any("args", args))
	}

	client := http.NewClient(cfg.HTTP, logger)

	var wg sync.WaitGroup
	for index := 0; index < count; index++ {
		go client.Request(&wg, args[0])
		wg.Add(1)
	}

	wg.Wait()
	logger.Info("all request has been finished")
}
