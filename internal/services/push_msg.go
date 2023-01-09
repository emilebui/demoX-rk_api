package services

import (
	"github.com/emilebui/demoX-rk_api/internal/repositories"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
)

type Service struct {
	r      repositories.Repo
	logger *rkentry.LoggerEntry
}

func NewService(r repositories.Repo, logger *rkentry.LoggerEntry) *Service {
	return &Service{
		r:      r,
		logger: logger,
	}
}

func (s *Service) Process(msg string) error {
	s.logger.Info("Processing message", zap.String("message", msg))
	return s.r.PushMessage(msg)
}
