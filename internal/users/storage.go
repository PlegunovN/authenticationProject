package users

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type client struct {
	db             *sqlx.DB
	logger         *zap.SugaredLogger
	tokenSecretKey string
}

func (s client) createUser(ctx context.Context, users Users) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Rollback,  createUser error: %w", err)
		}
		tx.Commit()
		return
	}()

	//записываем логин и хеш пароля
	query := "INSERT INTO users(login, password) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, users.Login, users.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s client) deleteUser(ctx context.Context, login string) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Delete DBUser error - rollback: %w", err)
			return
		}
		tx.Commit()
		return
	}()

	query := "DELETE FROM users  WHERE login = $1"
	_, err = tx.ExecContext(ctx, query, login)
	if err != nil {
		return err
	}

	return nil
}

func (s client) getUserPassword(ctx context.Context, login string) (string, error) {

	//получить hash по логину
	query := "SELECT password FROM users WHERE login = $1"
	var hashFromTable string
	err := s.db.GetContext(ctx, &hashFromTable, query, login)
	if err != nil {
		return "", nil
	}
	return hashFromTable, err
}
