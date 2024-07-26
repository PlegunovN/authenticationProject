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

func (s Service) SignUp(ctx context.Context, login, password, token string) error {
	err := s.client.signUp(ctx, Users{Login: login, Password: password}, Tokens{Token: token})
	return err
}

func (s Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.client.deleteUser(ctx, id)
	return err
}

func (s Service) SignIn(ctx context.Context, login, password string) (*Users, error) {
	user, err := s.client.signIn(ctx, login, password)
	return user, err
}
