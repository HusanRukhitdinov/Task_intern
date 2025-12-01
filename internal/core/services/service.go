package services

import (
	"intern/internal/core/repository/psql/sqlc"
	"intern/pkg/logger"
)

type Service struct {
	storage sqlc.Querier
	log     logger.ILogger
}

func NewService(storage sqlc.Querier, log logger.ILogger) *Service {
	return &Service{
		storage: storage,
		log:     log,
	}
}
