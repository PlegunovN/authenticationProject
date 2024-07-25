package users

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type client struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func (s client) signUp(ctx context.Context, users Users) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Rollback, sign up error: %w", err)
		}
		tx.Commit()
	}()
	query := "INSERT INTO users(login, password) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, users.Login, users.Password)
	if err != nil {
		return err
	}
	return err
}
