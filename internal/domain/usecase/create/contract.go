package create

import (
	"context"

	"github.com/xgmsx/go-tiny-url/internal/domain/entity"
)

//go:generate mockgen -source=contract.go -destination=mocks/contract.go

type database interface {
	CreateLink(context.Context, entity.Link) error
}

type cache interface {
	PutLink(context.Context, entity.Link) error
}

type publisher interface {
	SendLink(ctx context.Context, link entity.Link) error
}
