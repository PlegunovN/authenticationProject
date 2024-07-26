package handlers

import (
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Api struct {
	Storage *users.Service
	SLogger *zap.SugaredLogger
}

func New(db *sqlx.DB, logger *zap.SugaredLogger) *Api {
	return &Api{
		Storage: users.New(db, logger),
	}

}
