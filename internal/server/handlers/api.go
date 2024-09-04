package handlers

import (
	"github.com/PlegunovN/authenticationProject/internal/users"
	"go.uber.org/zap"
)

type Api struct {
	userService *users.Service
	logger      *zap.SugaredLogger
}

func New(service *users.Service, logger *zap.SugaredLogger) *Api {
	return &Api{
		userService: service,
		logger:      logger,
	}
}
