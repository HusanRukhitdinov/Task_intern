package repository

import (
	"context"
	"intern/internal/configs"
	"intern/internal/core/repository/psql"
	"intern/internal/core/repository/psql/sqlc"
	"intern/pkg/logger"
)

type StorageI interface {
	sqlc.Querier
}

func New(ctx context.Context, cfg configs.Config, log logger.ILogger) (StorageI, error) {
	return psql.NewStore(ctx, log, cfg)
}
