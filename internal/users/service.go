package users

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/PlegunovN/authenticationProject/internal/rabbit"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	client *client
}

func New(db *sqlx.DB, logger *zap.SugaredLogger, tokenSecretKey string) *Service {
	return &Service{
		client: &client{
			db:             db,
			logger:         logger,
			tokenSecretKey: tokenSecretKey,
		},
	}
}

// преобразование пароля в хэш
func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	hash := fmt.Sprint(sum)
	return hash
}

func jwtToken(tokenSecretKey []byte, login string) (string, error) {
	// Создаём данные для токена
	claims := make(jwt.MapClaims)
	claims["ExpiresAt"] = time.Now().Add(10 * time.Hour)
	claims["authorized"] = true
	claims["login"] = login

	// Генерируем подпись
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString(tokenSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil // Выводим сгенерированный токен
}

func (s Service) SignUp(ctx context.Context, login, password string) error {
	hash := hashPassword(password)
	id, err := s.client.createUser(ctx, Users{Login: login, Password: hash})
	if id != -1 {
		//передать токен в др сервис
		err = rabbit.Send(s.client.logger, login, id)
		if err != nil {
			s.client.logger.Errorf("send message error %w", err)
		}
	} else {
		return ErrorDuplicateLogin{}
	}
	return err
}

func (s Service) DeleteUser(ctx context.Context, login string) error {
	err := s.client.deleteUser(ctx, login)
	return err
}

func (s Service) SignIn(ctx context.Context, login, password string) (string, error) {
	hash := hashPassword(password)

	hashFromTable, err := s.client.getUserPassword(ctx, login)
	if hashFromTable == "" {
		return "", err
	}

	// сравнить хеш из базы и от пользователя
	if hashFromTable == hash {
		//создать токен jwt
		token, err := jwtToken([]byte(s.client.tokenSecretKey), login)
		if err != nil {
			return "", err
		}

		//передать токен юзеру
		return token, nil

	} else {
		return "", ErrorPasswordIncorrect{}
	}

}
