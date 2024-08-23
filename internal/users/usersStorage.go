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
			s.logger.Errorf("Delete User error - rollback: %w", err)
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

func (s client) getUserPasswordToValidate(ctx context.Context, login string) (string, error) {

	//получить hash по логину
	query := "SELECT password FROM users WHERE login = $1"
	var hashFromTable string
	err := s.db.GetContext(ctx, &hashFromTable, query, login)
	if err != nil {

		return "", nil
	}
	return hashFromTable, err
}

func (s client) generateUserToken(ctx context.Context, login, token string) (*Users, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Delete User error - rollback: %w", err)
			return
		}
		tx.Commit()
		return
	}()

	//записываем токен в базу
	query := "UPDATE users SET token = $1 WHERE login = $2"
	_, err = tx.ExecContext(ctx, query, token, login)
	if err != nil {
		return nil, err
	}

	// otдаем для примера user
	query = "SELECT * FROM users WHERE login = $1"
	user := &Users{}
	err = tx.GetContext(ctx, user, query, login)
	if err != nil {
		return nil, err
	}

	return user, err

}
