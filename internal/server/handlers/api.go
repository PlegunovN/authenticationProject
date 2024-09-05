package handlers

import (
	"github.com/PlegunovN/authenticationProject/configs"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"go.uber.org/zap"
)

type Api struct {
	userService    *users.Service
	logger         *zap.SugaredLogger
	tokenSecretKey *configs.SecretKey
}

func New(service *users.Service, logger *zap.SugaredLogger, tokenSecretKey *configs.SecretKey) *Api {
	return &Api{
		userService:    service,
		logger:         logger,
		tokenSecretKey: tokenSecretKey,
	}
}
