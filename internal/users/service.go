package users

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/PlegunovN/authenticationProject/configs"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	client *client
	logger *zap.SugaredLogger
}

func New(db *sqlx.DB, logger *zap.SugaredLogger) *Service {
	return &Service{
		client: &client{
			db:     db,
			logger: logger,
		},
	}
}

// преобразование пароля в хэш
func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	hash := fmt.Sprint(sum)
	return hash
}

func secretKey(logger zap.SugaredLogger) []byte {
	sk, err := configs.LoadSecretKey("./.env")
	if err != nil {
		logger.Error(err)
	}
	return []byte(sk.Key)
}

var TokenSecretKey = secretKey(zap.SugaredLogger{})

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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, nil
		}
		return TokenSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s Service) SignUp(ctx context.Context, login, password string) error {
	hash := hashPassword(password)
	err := s.client.createUser(ctx, Users{Login: login, Password: hash})
	return err
}

func (s Service) DeleteUser(ctx context.Context, login string) error {
	err := s.client.deleteUser(ctx, login)
	return err
}

func (s Service) SignIn(ctx context.Context, login, password string) (*Users, error) {
	hash := hashPassword(password)

	hashFromTable, err := s.client.getUserPasswordToValidate(ctx, login)
	if hashFromTable == "" {
		return nil, err
	}

	// сравнить хеш из базы и от пользователя
	if hashFromTable == hash {
		//создать токен jwt
		token, err := jwtToken(TokenSecretKey, login)
		//передать токен юзеру
		if err != nil {
			return nil, err
		}

		fmt.Println(token)
		return nil, nil

	} else {
		return nil, ErrorPasswordIncorrect{}
	}

}
