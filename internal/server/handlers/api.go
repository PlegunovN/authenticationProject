package handlers

import (
	"github.com/PlegunovN/authenticationProject/internal/users"
	"go.uber.org/zap"
)

type Api struct {
	userStorage *users.Service
	logger      *zap.SugaredLogger
}

func New(storage *users.Service, logger *zap.SugaredLogger) *Api {
	return &Api{
		userStorage: storage,
		logger:      logger,
	}
}
