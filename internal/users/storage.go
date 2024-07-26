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

func (s client) signUp(ctx context.Context, users Users, tokens Tokens) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			s.logger.Errorf("Rollback, sign up error: %w", err)
		}
		tx.Commit()
	}()

	//post login password
	query := "INSERT INTO users(login, password) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, users.Login, users.Password)
	if err != nil {
		return err
	}

	//берем только добавленные id and password
	query = "SELECT id, password FROM users WHERE login = $1 AND password = $2"
	u := new(Users)
	err = tx.GetContext(ctx, u, query, users.Login, users.Password)
	if err != nil {
		return err
	}

	//добавляем пароль в токен
	//обработать токен!
	query = "INSERT INTO tokens(token) VALUES ($1)"
	_, err = tx.ExecContext(ctx, query, u.Password)
	if err != nil {
		return err
	}

	////берем id из таблицы токенов
	query = "SELECT id FROM tokens WHERE token = $1"
	t := new(Tokens)
	err = tx.GetContext(ctx, t, query, tokens.Token)
	if err != nil {
		return err
	}
	//
	////соединяем таблицы юзеров и токенов
	query = "INSERT INTO usersandtokens(userid, tokenid) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, u.ID, t.ID)
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

	query := "SELECT tokens.id FROM tokens INNER JOIN usersandtokens ON tokens.id = usersandtokens.tokenid WHERE usersandtokens.userid = $1"
	var res int64
	err = tx.GetContext(ctx, res, query, id)

	query = "DELETE FROM usersandtokens  WHERE usersandtokens.userid = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM users WHERE users.id = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM tokens WHERE tokens.id=$1"
	_, err = tx.ExecContext(ctx, query, res)
	if err != nil {
		return err
	}

	return nil
}
