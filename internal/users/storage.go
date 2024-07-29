package users

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type client struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))

	hash := fmt.Sprint(sum)
	return hash
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

	//обработка хеш
	hash := hashPassword(users.Password)

	//записываем логин и хеш пароля
	query := "INSERT INTO users(login, password) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, users.Login, hash)
	if err != nil {
		return err
	}

	return nil
}

func (s client) deleteUser(ctx context.Context, id int64) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Delete User error - rollback: %w", err)
		}
		tx.Commit()
	}()

	query := "DELETE FROM users  WHERE id = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s client) signIn(ctx context.Context, login, paswFromUser string) (*Users, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("sign in error - rollback: %w", err)
		}
		tx.Commit()
	}()
	//получить токен по логину
	query := "SELECT password FROM users WHERE login = $1"
	var paswFromTable string
	err = tx.GetContext(ctx, &paswFromTable, query, login)
	if err != nil {
		return nil, err
	}

	//преобразовать пароль от пользователя в хеш
	pFU := hashPassword(paswFromUser)

	// сравнить хеш из базы и от пользователя
	if paswFromTable == pFU {
		query = "SELECT * FROM users WHERE login = $1 AND password = $2"
		user := &Users{}
		err = tx.GetContext(ctx, user, query, login, paswFromTable)
		if err != nil {
			return nil, err
		}

		return user, nil
	} else {
		incorectedPasW := "incorectedPassWord"
		user := &Users{Login: login, Password: incorectedPasW}
		return user, err
	}

	return nil, err
}
