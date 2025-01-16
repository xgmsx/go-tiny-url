package main

import (
	"context"

	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"

	"github.com/xgmsx/go-tiny-url/internal/app"
	"github.com/xgmsx/go-tiny-url/internal/config"
	"github.com/xgmsx/go-tiny-url/pkg/logger"
	"github.com/xgmsx/go-tiny-url/pkg/observability/otel"
	"github.com/xgmsx/go-tiny-url/pkg/observability/sentry"
)

var (
	cl configLoader = config.New()
	ar appRunner    = app.New()
)

func run(ctx context.Context, cl configLoader, ar appRunner) error {
	c, err := cl.Load(ctx)
	if err != nil {
		log.Error().Err(err).Msg("config.New")
		return err
	}

	logger.Init(c.Logger, c.App.Name, c.App.Version)
	log.Info().Msg("App starting...")
	defer log.Info().Msg("App stopped")

	sentry.Init(c.Sentry, c.App.Name, c.App.Version, c.App.Env)
	defer sentry.Close()

	otel.Init(ctx, c.Otel, c.App.Name, c.App.Version)
	defer otel.Close()

	err = ar.Run(ctx, c)
	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
	return err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx, cl, ar); err != nil {
		panic(err)
	}
}
