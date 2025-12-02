package services

import (
	"intern/internal/core/repository/psql/sqlc"
	"intern/pkg/email"
	"intern/pkg/logger"
)

type Service struct {
	storage     sqlc.Querier
	log         logger.ILogger
	emailSender *email.EmailSender
}

func NewService(storage sqlc.Querier, log logger.ILogger, emailSender *email.EmailSender) *Service {
	return &Service{
		storage:     storage,
		log:         log,
		emailSender: emailSender,
	}
}
