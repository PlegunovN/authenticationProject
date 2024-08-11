package users

import (
	"context"
	"errors"
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

func (s client) loginUser(ctx context.Context, login, hashFromUser, token string) (*Users, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("login error - rollback: %w", err)
			return
		}
		tx.Commit()
		return
	}()
	//получить токен по логину
	query := "SELECT password FROM users WHERE login = $1"
	var hashFromTable string
	err = tx.GetContext(ctx, &hashFromTable, query, login)
	if err != nil {
		return nil, err
	}

	// сравнить хеш из базы и от пользователя
	if hashFromTable == hashFromUser {

		//записываем токен в базу
		query = "UPDATE users SET token = $1 WHERE login = $2"
		_, err = tx.ExecContext(ctx, query, token, login)
		if err != nil {
			return nil, err
		}
		// otдаем для примера user
		query = "SELECT * FROM users WHERE login = $1 AND password = $2"
		user := &Users{}
		err = tx.GetContext(ctx, user, query, login, hashFromTable)
		if err != nil {
			return nil, err
		}

		return user, err

	} else {
		incorectedPasW := "incorectedPassWord"
		user := &Users{Login: login, Password: incorectedPasW}
		return user, err
	}

}

func (s client) work() (string, error) {

	return "it's working", nil
}

func (s client) validateToken(ctx context.Context, login, tokenFromUser string) error {
	//получаем токен из базы
	var tokenFromBase string
	query := "SELECT token FROM users WHERE login = $1"
	err := s.db.GetContext(ctx, &tokenFromBase, query, login)
	if err != nil {
		return err
	}
	if tokenFromBase != tokenFromUser {
		s.logger.Info("invalid token")
		err = errors.New("ощибка сравнения токенов")

		return err
	}

	return err
}
