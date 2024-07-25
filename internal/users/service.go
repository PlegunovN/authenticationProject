package users

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Service struct {
	client *client
}

func New(db *sqlx.DB, logger *zap.SugaredLogger) *Service {
	return &Service{
		client: &client{
			db:     db,
			logger: logger,
		},
	}
}

func (s Service) SignUp(ctx context.Context, login, password string) error {
	err := s.client.signUp(ctx, Users{Login: login, Password: password})
	return err
}
