package users

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
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

// преобразование пароля в хэш
func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	hash := fmt.Sprint(sum)
	return hash
}

func jwtToken(tokenSecretKey string) string {
	// Создаём данные для токена
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix() // Устанавливаем срок действия токена в 1 минут

	// Генерируем подпись
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(tokenSecretKey)) // Используем секретный ключ для подписи токена

	strToken := fmt.Sprint(signedToken)
	return strToken // Выводим сгенерированный токен
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

	//создать токен jwt
	tokenSecretKey := "secretKey"
	token := jwtToken(tokenSecretKey)

	user, err := s.client.loginUser(ctx, login, hash, token)
	return user, err
}

func (s Service) Work() (string, error) {

	resp, err := s.client.work()
	return resp, err
}

func (s Service) ValidateToken(ctx context.Context, login, token string) error {
	err := s.client.validateToken(ctx, login, token)
	return err
}
