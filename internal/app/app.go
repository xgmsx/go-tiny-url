package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/xgmsx/go-tiny-url/internal/config"
	adapterKafka "github.com/xgmsx/go-tiny-url/internal/domain/adapter/kafka"
	adapterPostgres "github.com/xgmsx/go-tiny-url/internal/domain/adapter/postgres"
	adapterRedis "github.com/xgmsx/go-tiny-url/internal/domain/adapter/redis"
	controllerGRPC "github.com/xgmsx/go-tiny-url/internal/domain/controller/grpc"
	controllerHTTP "github.com/xgmsx/go-tiny-url/internal/domain/controller/http"
	controllerKafka "github.com/xgmsx/go-tiny-url/internal/domain/controller/kafka"
	usecaseCreate "github.com/xgmsx/go-tiny-url/internal/domain/usecase/create"
	usecaseFetch "github.com/xgmsx/go-tiny-url/internal/domain/usecase/fetch"
	"github.com/xgmsx/go-tiny-url/pkg/grpc"
	"github.com/xgmsx/go-tiny-url/pkg/http"
	kafkaReader "github.com/xgmsx/go-tiny-url/pkg/kafka/reader"
	kafkaWriter "github.com/xgmsx/go-tiny-url/pkg/kafka/writer"
	postgresClient "github.com/xgmsx/go-tiny-url/pkg/postgres"
	redisClient "github.com/xgmsx/go-tiny-url/pkg/redis"
)

type App struct{}

func New() App {
	return App{}
}

func (a App) Run(ctx context.Context, c *config.Config) error {
	// init dependencies
	postgres, err := postgresClient.New(ctx, &c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer postgres.Close()

	redis, err := redisClient.New(&c.Redis)
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}
	defer redis.Close()

	KafkaWriter, err := kafkaWriter.New(&c.KafkaWriter)
	if err != nil {
		return fmt.Errorf("kafkaWriter.New: %w", err)
	}
	defer KafkaWriter.Close()

	KafkaReader, err := kafkaReader.New(&c.KafkaReader)
	if err != nil {
		return fmt.Errorf("kafkaReader.New: %w", err)
	}
	defer KafkaReader.Close()

	// init adapter
	database := adapterPostgres.New(postgres.Pool)
	cache := adapterRedis.New(redis.Client)
	publisher := adapterKafka.New(KafkaWriter.Writer)

	// init usecase
	ucCreateLink := usecaseCreate.New(database, cache, publisher)
	ucFetchLink := usecaseFetch.New(database, cache)

	// init controller
	errCh := make(chan error)
	defer close(errCh)

	httpServer := http.New(c.HTTP, nil, controllerHTTP.New("/api/shortener", ucCreateLink, ucFetchLink))
	go func() { errCh <- httpServer.Serve(c.HTTP.Port) }()
	defer httpServer.Close()

	grpcServer := grpc.New(controllerGRPC.New(ucCreateLink, ucFetchLink))
	go func() { errCh <- grpcServer.Serve(ctx, c.GRPC.Port) }()
	defer grpcServer.Close()

	kafkaConsumer := controllerKafka.New(KafkaReader, ucCreateLink)
	go func() { errCh <- kafkaConsumer.Consume(ctx) }()

	return a.waiting(errCh)
}

func (a App) waiting(errCh <-chan error) error {
	log.Info().Msg("App started")
	defer log.Info().Msg("App stopping...")

	ctxTerm, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	select {
	case <-ctxTerm.Done():
		log.Info().Msg("App got termination signal")
		return nil
	case err := <-errCh:
		log.Info().Err(err).Msg("App got error notify")
		return err
	}
}
