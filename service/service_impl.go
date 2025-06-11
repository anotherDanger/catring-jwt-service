package service

import (
	"catering-jwt-service/domain"
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ServiceImpl struct {
}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}

func (svc *ServiceImpl) Register(ctx context.Context, entity *domain.Domain) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       entity.Id,
		"username": entity.Username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})

	tokenT := os.Getenv("JWT_SECRET")
	byteToken := []byte(tokenT)

	token, err := claims.SignedString(byteToken)
	if err != nil {
		panic(err)
	}

	return token, nil
}
