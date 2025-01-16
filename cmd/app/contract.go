package main

import (
	"context"

	"github.com/xgmsx/go-tiny-url/internal/config"
)

//go:generate mockgen -source=contract.go -destination=mocks/contract.go

type appRunner interface {
	Run(ctx context.Context, c *config.Config) error
}

type configLoader interface {
	Load(ctx context.Context) (*config.Config, error)
}
