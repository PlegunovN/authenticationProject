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
	err := s.client.createUser(ctx, Users{Login: login, Password: password})
	return err
}

func (s Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.client.deleteUser(ctx, id)
	return err
}

func (s Service) SignIn(ctx context.Context, login, password string) (*Users, string, error) {
	user, token, err := s.client.loginUser(ctx, login, password)
	return user, token, err
}

func (s Service) Work(ctx context.Context, login, tokenFromUser string) (string, error) {
	resp, err := s.client.work(ctx, login, tokenFromUser)
	return resp, err
}
