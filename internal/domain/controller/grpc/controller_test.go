package grpc

import (
	"context"
	"testing"

	"github.com/xgmsx/go-tiny-url/internal/domain/usecase/create"
	"github.com/xgmsx/go-tiny-url/internal/domain/usecase/fetch"
	"github.com/xgmsx/go-tiny-url/pkg/grpc"
)

func TestController(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := New(create.Usecase{}, fetch.Usecase{})
	srv := grpc.New(ctrl)
	defer srv.Close()

	go func() {
		_ = srv.Serve(ctx, "9090")
	}()
}
