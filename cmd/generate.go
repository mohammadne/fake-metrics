package cmd

import (
	"os"

	"github.com/MohammadNE/fake-metrics/internal/config"
	"github.com/MohammadNE/fake-metrics/pkg/logger"
	"github.com/spf13/cobra"
)

type Generate struct{}

func (m Generate) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.main(config.Load(true), args, trap)
	}

	return &cobra.Command{
		Use:       "generate",
		Short:     "generate (simulate) HTTP requests",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m *Generate) main(cfg *config.Config, args []string, trap chan os.Signal) {
	logger := logger.NewZap(cfg.Logger)
	_ = logger

	// if len(args) != 1 {
	// 	logger.Fatal("invalid arguments given", zap.Any("args", args))
	// }

	// logger.Info("Database has been migrated successfully", zap.String("migration", args[0]))
}
